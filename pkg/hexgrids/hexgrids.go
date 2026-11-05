package hexgrids

import "github.com/PhilAldridge/aoc-2024-go/pkg/coords"

type HexGrid [][]int

func (hexGrid HexGrid) DownRightRows() [][]coords.Coord {
	out := [][]coords.Coord{}
	for i := range hexGrid[0] {
		out = append(out, hexGrid.getRow(coords.NewCoord(0, i), downRight))
	}

	for i := 1; i <= len(hexGrid)/2; i++ {
		out = append(out, hexGrid.getRow(coords.NewCoord(i, 0), downRight))
	}
	return out
}

func (hexGrid HexGrid) DownLeftRows() [][]coords.Coord {
	out := [][]coords.Coord{}
	for i := range hexGrid[0] {
		out = append(out, hexGrid.getRow(coords.NewCoord(0, i), downLeft))
	}

	for i := 1; i <= len(hexGrid)/2; i++ {
		out = append(out, hexGrid.getRow(coords.NewCoord(i, len(hexGrid[i])-1), downLeft))
	}
	return out
}

func (hexGrid HexGrid) LeftRows() [][]coords.Coord {
	out := [][]coords.Coord{}
	for i := range hexGrid {
		out = append(out, hexGrid.getRow(coords.NewCoord(i, len(hexGrid[i])-1), left))
	}
	return out
}

func (hexGrid HexGrid) UpLeftRows() [][]coords.Coord {
	out := [][]coords.Coord{}
	for i := range hexGrid[0] {
		out = append(out, hexGrid.getRow(coords.NewCoord(len(hexGrid)-1, i), upLeft))
	}

	for i := len(hexGrid) / 2; i < len(hexGrid)-1; i++ {
		out = append(out, hexGrid.getRow(coords.NewCoord(i, len(hexGrid[i])-1), upLeft))
	}
	return out
}

func (hexGrid HexGrid) UpRightRows() [][]coords.Coord {
	out := [][]coords.Coord{}
	for i := range hexGrid[0] {
		out = append(out, hexGrid.getRow(coords.NewCoord(len(hexGrid)-1, i), upRight))
	}

	for i := len(hexGrid) / 2; i < len(hexGrid)-1; i++ {
		out = append(out, hexGrid.getRow(coords.NewCoord(i, 0), upRight))
	}
	return out
}

func (hexGrid HexGrid) RightRows() [][]coords.Coord {
	out := [][]coords.Coord{}
	for i := range hexGrid {
		out = append(out, hexGrid.getRow(coords.NewCoord(i, 0), right))
	}
	return out
}

func (hexGrid HexGrid) GetAllDirectionRows() [][][]coords.Coord {
	return [][][]coords.Coord{
		hexGrid.DownRightRows(),
		hexGrid.DownLeftRows(),
		hexGrid.LeftRows(),
		hexGrid.UpLeftRows(),
		hexGrid.UpRightRows(),
		hexGrid.RightRows(),
	}
}

func downRight(pos coords.Coord, sideLength int) (coords.Coord, bool) {
	newI := pos.I + 1
	if newI+1 >= sideLength*2 {
		return coords.Coord{}, false
	}

	if newI < sideLength {
		return coords.NewCoord(newI, pos.J+1), true
	}

	length := sideLength*3 - newI - 2
	if pos.J >= length {
		return coords.Coord{}, false
	}

	return coords.NewCoord(newI, pos.J), true
}

func downLeft(pos coords.Coord, sideLength int) (coords.Coord, bool) {
	newI := pos.I + 1
	if newI+1 >= sideLength*2 {
		return coords.Coord{}, false
	}

	if newI < sideLength {
		return coords.NewCoord(newI, pos.J), true
	}
	if pos.J < 1 {
		return coords.Coord{}, false
	}

	return coords.NewCoord(newI, pos.J-1), true
}

func left(pos coords.Coord, sideLength int) (coords.Coord, bool) {
	if pos.J < 1 {
		return coords.Coord{}, false
	}

	return coords.NewCoord(pos.I, pos.J-1), true
}

func upLeft(pos coords.Coord, sideLength int) (coords.Coord, bool) {
	newI := pos.I - 1
	if newI < 0 {
		return coords.Coord{}, false
	}

	if newI < sideLength-1 {
		if pos.J < 1 {
			return coords.Coord{}, false
		}

		return coords.NewCoord(newI, pos.J-1), true
	}

	return coords.NewCoord(newI, pos.J), true
}

func upRight(pos coords.Coord, sideLength int) (coords.Coord, bool) {
	newI := pos.I - 1
	if newI < 0 {
		return coords.Coord{}, false
	}

	if newI < sideLength-1 {
		length := sideLength + newI
		if pos.J >= length {
			return coords.Coord{}, false
		}

		return coords.NewCoord(newI, pos.J), true
	}

	return coords.NewCoord(newI, pos.J+1), true
}

func right(pos coords.Coord, sideLength int) (coords.Coord, bool) {
	length := sideLength + pos.I
	if pos.I >= sideLength {
		length = sideLength*3 - pos.I - 2
	}
	if pos.J+1 >= length {
		return coords.Coord{}, false
	}

	return coords.NewCoord(pos.I, pos.J+1), true
}

func (hexGrid HexGrid) getRow(start coords.Coord, moveFun func(pos coords.Coord, sideLength int) (coords.Coord, bool)) []coords.Coord {
	ok := true
	row := []coords.Coord{start}
	pos := start

	for {
		pos, ok = moveFun(pos, hexGrid.SideLength())
		if !ok {
			break
		}

		row = append(row, pos)
	}

	return row
}

func (hexGrid HexGrid) SideLength() int {
	return len(hexGrid[0])
}
