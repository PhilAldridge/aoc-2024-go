package complexn

type Complex struct {
	R,I int
}

func NewComplex(r,i int) Complex {
	return Complex{
		R:r,
		I: i,
	}
}

func Add(a,b Complex) Complex {
	return Complex{
		R:a.R + b.R,
		I:a.I + b.I,
	}
}

func Subtract(a,b Complex) Complex {
	return Complex{
		R:a.R - b.R,
		I:a.I - b.I,
	}
}

func Multiply(a,b Complex) Complex {
	return Complex{
		R:a.R*b.R - a.I*b.I,
		I:a.R*b.I + a.I*b.R,
	}
}

func DivideBasic(a,b Complex) Complex {
	return Complex{
		R:a.R/b.R,
		I:a.I/b.I,
	}
}