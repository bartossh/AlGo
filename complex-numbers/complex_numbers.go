package complexnumbers

import "math"

// Number is a complex number entity
type Number struct {
	real float64
	imag float64
}

func (n Number) Real() float64 {
	return n.real
}

func (n Number) Imaginary() float64 {
	return n.imag
}

func (n1 Number) Add(n2 Number) Number {
	real := (n1.real + n2.real)
	imag := (n1.imag + n2.imag)
	return Number{real, imag}
}

func (n1 Number) Subtract(n2 Number) Number {
	real := (n1.real - n2.real)
	imag := (n1.imag - n2.imag)
	return Number{real, imag}
}

func (n1 Number) Multiply(n2 Number) Number {
	real := n1.real*n2.real - n1.imag*n2.imag
	imag := n1.imag*n2.real + n1.real*n2.imag
	return Number{real, imag}
}

func (n Number) Times(factor float64) Number {
	return Number{n.real * factor, n.imag * factor}
}

func (n1 Number) Divide(n2 Number) Number {
	real := (n1.real*n2.real + n1.imag*n2.imag) / (math.Pow(n2.real, 2) + math.Pow(n2.imag, 2))
	imag := (n1.imag*n2.real - n1.real*n2.imag) / (math.Pow(n2.real, 2) + math.Pow(n2.imag, 2))
	return Number{real, imag}
}

func (n Number) Conjugate() Number {
	return Number{n.real, -n.imag}
}

func (n Number) Abs() float64 {
	return math.Sqrt(math.Pow(n.real, 2) + math.Pow(n.imag, 2))
}

func (n Number) Exp() Number {
	real := math.Exp(n.real) * math.Cos(n.imag)
	imag := math.Exp(n.real) * math.Sin(n.imag)
	return Number{real, imag}
}
