package graphs

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"slices"
	"strings"
	"sync"
)

// Node mirrors your struct: Energy drives the node's size, Connections are
// outgoing (directed) edges to other nodes.
type Node struct {
	Size      int
	Connections []Connection
}

// Connection is one outgoing edge. Weight drives the edge's thickness.
type Connection struct {
	To     int
	Weight int
}

type Graph map[int]Node

// Size mapping: area grows with energy, so width grows with sqrt(energy).
const (
	minWidth   = 0.3  // inches, for energy 0
	widthScale = 0.15 // inches per sqrt(energy)
)

func nodeWidth(energy int) float64 {
	return widthScale*math.Sqrt(float64(max(energy, 0)))/300.0
}

// Edge thickness: line width grows linearly with weight.
const (
	minPenWidth = 0.5 // points, for weight 0
	penScale    = 0.1 // points per unit of weight
)

func edgePenWidth(weight int) float64 {
	return minPenWidth + penScale*float64(max(weight, 0))
}

// DOT renders the graph deterministically (sorted IDs) so successive frames
// diff cleanly. Every node and edge gets a stable `id` attribute, which
// d3-graphviz uses (keyMode "id") to match elements between frames and
// animate them rather than redrawing from scratch.
func (g Graph) DOT() string {
	ids := make([]int, 0, len(g))
	for id := range g {
		ids = append(ids, id)
	}
	slices.Sort(ids)

	var b strings.Builder
	b.WriteString("digraph G {\n")
	b.WriteString("  graph [bgcolor=\"transparent\", splines=false, scale=prism, sep=\"+1200\", K=3];\n")
	b.WriteString("  node [shape=circle, fixedsize=true, style=filled, fillcolor=\"#cfe3f5\", color=\"#2b6ca3\", fontname=\"Helvetica\", fontsize=11];\n")
	b.WriteString("  edge [color=\"#7a8794\", arrowsize=0.7];\n")

	for _, id := range ids {
		n := g[id]
		w := nodeWidth(n.Size)
		fmt.Fprintf(&b, "  n%d [id=\"n%d\", label=\"%d\", width=%.3f, height=%.3f, tooltip=\"node %d, energy %d\"];\n",
			id, id, id, w, w, id, n.Size)
	}

	for _, from := range ids {
		conns := slices.Clone(g[from].Connections)
		slices.SortStableFunc(conns, func(a, b Connection) int { return a.To - b.To })
		// duplicate edges would share an id; the first one for each target wins
		conns = slices.CompactFunc(conns, func(a, b Connection) bool { return a.To == b.To })
		for _, c := range conns {
			if _, ok := g[c.To]; !ok {
				continue // DOT would otherwise invent a phantom node
			}
			pw := edgePenWidth(c.Weight)
			// grow the arrowhead a little with the line so thick edges don't end in a tiny arrow
			arrow := 0.5 + 0.15*math.Sqrt(pw)
			fmt.Fprintf(&b, "  n%d -> n%d [id=\"e%d_%d\", penwidth=%.2f, arrowsize=%.2f, tooltip=\"%d -> %d, weight %d\"];\n",
				from, c.To, from, c.To, pw, arrow, from, c.To, c.Weight)
		}
	}
	b.WriteString("}\n")
	return b.String()
}

// Hub owns the graph and broadcasts a fresh DOT frame to every connected
// browser after each change. Slow clients only ever get the latest frame.
type Hub struct {
	mu    sync.Mutex
	graph Graph
	dot   string
	subs  map[chan string]struct{}
}

func NewHub(g Graph) *Hub {
	return &Hub{graph: g, dot: g.DOT(), subs: map[chan string]struct{}{}}
}

// Update applies fn to the graph under the lock, then pushes the new frame.
func (h *Hub) Update(fn func(Graph)) {
	h.mu.Lock()
	defer h.mu.Unlock()
	fn(h.graph)
	h.dot = h.graph.DOT()
	for ch := range h.subs {
		sendLatest(ch, h.dot)
	}
}

func (h *Hub) subscribe() (chan string, func()) {
	ch := make(chan string, 1)
	h.mu.Lock()
	h.subs[ch] = struct{}{}
	ch <- h.dot // current state straight away
	h.mu.Unlock()
	return ch, func() {
		h.mu.Lock()
		delete(h.subs, ch)
		h.mu.Unlock()
	}
}

func sendLatest(ch chan string, v string) {
	select {
	case ch <- v:
	default:
		select { // drop the stale frame, keep the newest
		case <-ch:
		default:
		}
		ch <- v
	}
}

// ServeHTTP streams DOT frames as Server-Sent Events.
// With Echo: e.GET("/events", echo.WrapHandler(hub)).
func (h *Hub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ch, unsubscribe := h.subscribe()
	defer unsubscribe()

	for {
		select {
		case <-r.Context().Done():
			return
		case dot := <-ch:
			payload, _ := json.Marshal(dot) // one line, safe for SSE
			fmt.Fprintf(w, "data: %s\n\n", payload)
			flusher.Flush()
		}
	}
}