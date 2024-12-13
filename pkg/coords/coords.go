package coords

import "github.com/PhilAldridge/aoc-2024-go/pkg/ints"

type Coord struct {
	I int
	J int
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

func MovementVector(a Coord, b Coord) Coord {
	iVec := a.I - b.I
	jVec := a.J - b.J
	gcd :=  ints.GCD(iVec,jVec)
	return Coord{
		I: iVec / gcd,
		J: jVec / gcd,
	}
}

func (a Coord) Add (b Coord) Coord {
	return Coord{
		I: a.I+b.I,
		J: a.J+b.J,
	}
}

func (a Coord) Subtract (b Coord) Coord {
	return Coord{
		I: a.I-b.I,
		J: a.J-b.J,
	}
}

func (a Coord) OnLine (c Coord, m Coord) bool {
	return (a.I - c.I)*m.J == (a.J-c.J)*m.I
}

func (a Coord) MoveBy (m Coord, times int) Coord {
	return Coord{
		I: a.I + m.I*times,
		J: a.J + m.J*times,
	}
}

func NewCoord (i int, j int) Coord {
	return Coord{
		I:i, 
		J:j,
	}
}

func (a Coord) SameDirectionAs (b Coord) bool {
	if (a.I ==0 && a.J == 0) || (b.I==0 && b.J==0) {
		return false
	}
	return a.I * b.J == a.J * b.I
}