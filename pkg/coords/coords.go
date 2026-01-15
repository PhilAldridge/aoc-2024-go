package coords

import "github.com/PhilAldridge/aoc-2024-go/pkg/ints"

type Coord struct {
	I int
	J int
}

type Line struct {
	M Coord
	C Coord
}

func (a Coord) Up(amount int) Coord {
	return Coord{
		I: a.I - amount,
		J: a.J,
	}
}

func (a Coord) Down(amount int) Coord {
	return Coord{
		I: a.I + amount,
		J: a.J,
	}
}

func (a Coord) Left(amount int) Coord {
	return Coord{
		I: a.I,
		J: a.J - amount,
	}
}

func (a Coord) Right(amount int) Coord {
	return Coord{
		I: a.I,
		J: a.J + amount,
	}
}

func (a Coord) GetAdjacent() [4]Coord {
	return [4]Coord{
		a.Up(1),
		a.Down(1),
		a.Left(1),
		a.Right(1),
	}
}

func (a Coord) IsOposite(b Coord) bool {
	return a.I == -b.I && a.J == -b.J
}

func (a Coord) GetAdjacentIncludingDiagonals() [8]Coord {
	return [8]Coord{
		a.Up(1),
		a.Up(1).Right(1),
		a.Right(1),
		a.Down(1).Right(1),
		a.Down(1),
		a.Down(1).Left(1),
		a.Left(1),
		a.Up(1).Left(1),
	}
}

func (a Coord) GetAdjacentDiagonals() [4]Coord {
	return [4]Coord{
		a.Up(1).Right(1),
		a.Down(1).Right(1),
		a.Down(1).Left(1),
		a.Up(1).Left(1),
	}
}

func MovementVector(a Coord, b Coord) Coord {
	iVec := a.I - b.I
	jVec := a.J - b.J
	gcd := ints.GCD(iVec, jVec)
	return Coord{
		I: iVec / gcd,
		J: jVec / gcd,
	}
}

func (a Coord) Add(b Coord) Coord {
	return Coord{
		I: a.I + b.I,
		J: a.J + b.J,
	}
}

func (a Coord) Subtract(b Coord) Coord {
	return Coord{
		I: a.I - b.I,
		J: a.J - b.J,
	}
}

func (a Coord) IsOnLine(l Line) bool {
	return (a.I-l.C.I)*l.M.J == (a.J-l.C.J)*l.M.I
}

func (a Coord) MoveBy(m Coord, times int) Coord {
	return Coord{
		I: a.I + m.I*times,
		J: a.J + m.J*times,
	}
}

func NewCoord(i int, j int) Coord {
	return Coord{
		I: i,
		J: j,
	}
}

func (a Coord) IsSameDirectionAs(b Coord) bool {
	if (a.I == 0 && a.J == 0) || (b.I == 0 && b.J == 0) {
		return false
	}
	return a.I*b.J == a.J*b.I
}

func LinesIntersect(l1 Line, l2 Line) bool {
	if l1.M.IsSameDirectionAs(l2.M) {
		return l1.C.IsOnLine(l2)
	}
	return true
}

func IntersectionPoint(l1 Line, l2 Line) Coord {
	if l1.M.IsSameDirectionAs(l2.M) {
		panic("intersectionPoint function not designed for parallel lines")
	}
	l1D := dotProduct(l1.M, l1.C) * 1.0
	l2D := dotProduct(l2.M, l2.C) * 1.0
	det := (l1.M.I*l2.M.J - l1.M.J*l2.M.I) * 1.0
	return NewCoord(
		(l1D*l2.M.J-l2D*l2.M.I)/det,
		(l1.M.I*l2D-l1.M.J*l1D)/det,
	)
}

func dotProduct(a Coord, b Coord) int {
	return a.I*b.I + a.J*b.J
}

func NewLine(m Coord, c Coord) Line {
	return Line{
		M: m,
		C: c,
	}
}

func (a Coord) Multiply(lambda int) Coord {
	return Coord{
		I: a.I * lambda,
		J: a.J * lambda,
	}
}

func CoordInSlice(c Coord, sl []Coord) bool {
	for _, s := range sl {
		if c.I == s.I && c.J == s.J {
			return true
		}
	}
	return false
}

var DirectionsInOrder = [4]Coord{
	NewCoord(0, 1),
	NewCoord(-1, 0),
	NewCoord(0, -1),
	NewCoord(1, 0),
}

func TurnLeft(a Coord) Coord {
	for i, v := range DirectionsInOrder {
		if v.Equals(a) {
			return DirectionsInOrder[(i+1)%4]
		}
	}
	panic("")
}

func TurnRight(a Coord) Coord {
	for i, v := range DirectionsInOrder {
		if v.Equals(a) {
			return DirectionsInOrder[(i+3)%4]
		}
	}
	panic("")
}

func TurnBack(a Coord) Coord {
	for i, v := range DirectionsInOrder {
		if v.IsSameDirectionAs(a) {
			return DirectionsInOrder[(i+2)%4]
		}
	}
	panic("")
}

func ManhattanDistance(a Coord, b Coord) int {
	x := a.J - b.J
	y := a.I - b.I
	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}
	return x + y
}

func PythagoreanSquareDistance(a,b Coord) int {
	iDiff:= a.I-b.I
	jDiff:= a.J-b.J

	return iDiff*iDiff + jDiff*jDiff
}

func (a Coord) Equals(b Coord) bool {
	return a.I == b.I && a.J == b.J
}

func (a Coord) InInput(input []string) bool {
	if a.I < 0 || a.J < 0 || a.I >= len(input) || a.J >= len(input[0]) {
		return false
	}
	return true
}

type Lenable interface {
	~string |
	~[]byte |
	~[]bool |
	~[]rune |
	~[]int |
	~[]any |
	~map[string]int |
	~map[string]string
}

func GenericInInput[T Lenable](a Coord, input []T) bool {
		if a.I < 0 || a.J < 0 || a.I >= len(input) || a.J >= len(input[0]) {
		return false
	}
	return true
}

func (a Coord) GetKnightMoves() []Coord {
	return []Coord{
		a.Add(NewCoord(1,2)),
		a.Add(NewCoord(1,-2)),
		a.Add(NewCoord(-1,2)),
		a.Add(NewCoord(-1,-2)),
		a.Add(NewCoord(2,1)),
		a.Add(NewCoord(2,-1)),
		a.Add(NewCoord(-2,1)),
		a.Add(NewCoord(-2,-1)),
	}
}
