package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/ints"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1 answer: ", part1("input.txt"))
	split := time.Now()
	fmt.Println("Part 2 answer: ", part2("input2.txt"))
	split2 := time.Now()
	fmt.Println("Part 3 answer: ", part3("input3.txt"))

	fmt.Println()
	fmt.Println("Part 1: ", split.Sub(start))
	fmt.Println("Part 2: ", split2.Sub(split))
	fmt.Println("Part 2: ", time.Since(split2))
}

func part1(name string) int {
	nodes := parseInput(name)

	root := connectNodes(nodes, true)

	readNodes := root.read()

	total := 0

	for i, node := range readNodes {
		total += (i + 1) * node.id
	}

	return total
}

func part2(name string) int {
	nodes := parseInput(name)

	root := connectNodes(nodes, false)

	readNodes := root.read()

	total := 0

	for i, node := range readNodes {
		total += (i + 1) * node.id
	}

	return total
}

func part3(name string) int {
	nodes := parseInput(name)

	root := connectNodesWithBreaking(nodes)

	readNodes := root.read()

	total := 0

	for i, node := range readNodes {
		total += (i + 1) * node.id
		fmt.Println(node.data)
	}

	return total
}

type nodeType struct {
	id                            int
	plug, leftSocket, rightSocket string
	left, right                   *nodeType
	data                          string
}

func parseInput(name string) []*nodeType {
	lines := files.ReadLines(name)

	nodes := []*nodeType{}
	for _, line := range lines {
		node := parseNode(line)
		nodes = append(nodes, &node)
	}

	return nodes
}

func parseNode(input string) nodeType {
	split := strings.Split(input, ", ")
	return nodeType{
		id:          ints.FromString(afterEqual(split[0])),
		plug:        afterEqual(split[1]),
		leftSocket:  afterEqual(split[2]),
		rightSocket: afterEqual(split[3]),
		data:        afterEqual(split[4]),
	}
}

func connectNodes(input []*nodeType, onlyStrong bool) *nodeType {
	root := input[0]

	for _, other := range input[1:] {
		connectNode(other, root, onlyStrong)
	}

	return root
}

func connectNode(node *nodeType, root *nodeType, onlyStrong bool) bool {
	if root.left == nil && canConnect(root.leftSocket, node.plug, onlyStrong) {
		root.left = node
		return true
	}

	if root.left != nil && connectNode(node, root.left, onlyStrong) {
		return true
	}

	if root.right == nil && canConnect(root.rightSocket, node.plug, onlyStrong) {
		root.right = node
		return true
	}

	if root.right == nil {
		return false
	}

	return connectNode(node, root.right, onlyStrong)
}

func connectNodesWithBreaking(input []*nodeType) *nodeType {
	root := input[0]

	for _, other := range input[1:] {
		for other != nil {
			other = connectNodeWithBreaking(other, root)
		}
	}

	return root
}

func connectNodeWithBreaking(node, root *nodeType) *nodeType {
	if root.left == nil {
		if canConnect(root.leftSocket, node.plug, false) {
			root.left = node
			return nil
		}
	} else {
		if root.left.plug != root.leftSocket && root.leftSocket == node.plug {
			root.left, node = node, root.left
		} else if node = connectNodeWithBreaking(node, root.left); node == nil {
			return nil
		}
	}

	if root.right == nil {
		if canConnect(root.rightSocket, node.plug, false) {
			root.right = node
			return nil
		}

		return node
	}

	if root.right.plug != root.rightSocket && root.rightSocket == node.plug {
		root.right, node = node, root.right

		return node
	}

	return connectNodeWithBreaking(node, root.right)
}

func canConnect(a, b string, onlyStrong bool) bool {
	if onlyStrong {
		return a == b
	}

	splita := strings.Split(a, " ")
	splitb := strings.Split(b, " ")

	return splita[0] == splitb[0] || splita[1] == splitb[1]
}

func afterEqual(input string) string {
	i := strings.Index(input, "=")
	return input[i+1:]
}

func (n *nodeType) read() []nodeType {
	if n == nil {
		return []nodeType{}
	}

	out := []nodeType{}

	out = append(out, n.left.read()...)
	out = append(out, *n)
	out = append(out, n.right.read()...)

	return out
}
