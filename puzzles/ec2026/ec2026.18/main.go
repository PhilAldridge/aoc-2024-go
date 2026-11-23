package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
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
	total := 0

	stackPositions := parse(name)

	lCursor, rCursor := 0, len(stackPositions)-1
	lastHeight := 0
	currentHeight, rect := getRect(stackPositions[lCursor], stackPositions[rCursor], lastHeight)
	total += rect

	for lCursor < rCursor {
		for stackPositions[lCursor].y <= currentHeight && lCursor < rCursor {
			lCursor++
		}

		for stackPositions[rCursor].y <= currentHeight && lCursor < rCursor {
			rCursor--
		}

		lastHeight = currentHeight
		currentHeight, rect = getRect(stackPositions[lCursor], stackPositions[rCursor], lastHeight)
		total += rect
	}

	return total
}

func part2(name string) int {
	stackPositions := parse(name)

	renderImage(stackPositions, "test.png", 5000, 1000)

	return part1(name)
}

func part3(name string) int {
	return part1(name)
}

type keyType struct {
	x, y  int
	width int
}

func newKey(vals []int) keyType {
	if len(vals) < 3 {
		vals = append(vals, 1)
	}

	return keyType{
		x: vals[0], y: vals[1], width: vals[2],
	}
}

func parse(name string) []keyType {
	out := []keyType{}

	for _, row := range files.ReadLines(name) {
		vals := ints.FromDelimitedString(row, " ")

		out = append(out, newKey(vals))
	}

	return out
}

func getRect(l, r keyType, prevHeight int) (newHeight int, rect int) {
	if l.y < r.y {
		newHeight = l.y
	} else {
		newHeight = r.y
	}

	if newHeight < prevHeight {
		panic("height went down")
	}

	rect = (r.x - l.x + r.width) * (newHeight - prevHeight)

	return newHeight, rect
}

func renderImage(walls []keyType, filename string, imgW, imgH int) error {
	n := len(walls)
	maxX := walls[n-1].x + walls[n-1].width

	prefixMax := make([]int, n)
	suffixMax := make([]int, n)
	for i := range n {
		prefixMax[i] = walls[i].y
		if i > 0 {
			prefixMax[i] = max(prefixMax[i], prefixMax[i-1])
		}
	}
	for i := n - 1; i >= 0; i-- {
		suffixMax[i] = walls[i].y
		if i < n-1 {
			suffixMax[i] = max(suffixMax[i], suffixMax[i+1])
		}
	}
	maxY := prefixMax[n-1]

	// Separate scales so the image is filled; use one scale for true proportions.
	sx := float64(imgW) / float64(maxX)
	sy := float64(imgH) / float64(maxY)

	img := image.NewRGBA(image.Rect(0, 0, imgW, imgH))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	fill := func(x0, x1, y0, y1 int, c color.Color) {
		px0 := int(float64(x0) * sx)
		px1 := max(px0+1, int(float64(x1)*sx+0.999)) // at least 1px wide
		py1 := imgH - int(float64(y0)*sy)            // flip: y=0 at the bottom
		py0 := imgH - int(float64(y1)*sy+0.999)
		if py0 >= py1 {
			return
		}
		r := image.Rect(px0, py0, px1, py1).Intersect(img.Bounds())
		draw.Draw(img, r, &image.Uniform{c}, image.Point{}, draw.Src)
	}

	blue := color.RGBA{40, 110, 220, 255}
	green := color.RGBA{30, 160, 60, 255}

	// Water first, so walls are drawn on top.
	for i, w := range walls {
		level := min(prefixMax[i], suffixMax[i])
		fill(w.x, w.x+w.width, w.y, level, blue) // water sitting on top of a wall
		if i+1 < n {
			gapLevel := min(prefixMax[i], suffixMax[i+1])
			fill(w.x+w.width, walls[i+1].x, 0, gapLevel, blue) // water between walls
		}
	}
	for _, w := range walls {
		fill(w.x, w.x+w.width, 0, w.y, green)
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, img)
}
