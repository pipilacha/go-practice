package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
	a float64
	b float64
}

func (n Number) Real() float64 {
	return n.a
}

func (n Number) Imaginary() float64 {
	return n.b
}

func (n1 Number) Add(n2 Number) Number {
	n1.a += n2.a
	n1.b += n2.b
	return n1
}

func (n1 Number) Subtract(n2 Number) Number {
	n1.a -= n2.a
	n1.b -= n2.b
	return n1
}

func (n1 Number) Multiply(n2 Number) Number {
	na, nb := n1.a, n1.b
	n1.a = na*n2.a - nb*n2.b
	n1.b = na*n2.b + nb*n2.a
	return n1
}

func (n Number) Times(factor float64) Number {
	n.a *= factor
	n.b *= factor
	return n
}

func (n1 Number) Divide(n2 Number) Number {
	na, nb := n1.a, n1.b
	n1.a = (na*n2.a + nb*n2.b) / (math.Pow(n2.a, 2) + math.Pow(n2.b, 2))
	n1.b = (nb*n2.a - na*n2.b) / (math.Pow(n2.a, 2) + math.Pow(n2.b, 2))
	return n1
}

func (n Number) Conjugate() Number {
	n.b *= -1
	return n
}

func (n Number) Abs() float64 {
	return math.Sqrt(math.Pow(n.a, 2) + math.Pow(n.b, 2))
}

func (n Number) Exp() Number {
	a, b := n.a, n.b
	n.a = math.Pow(math.E, a) * math.Cos(b)
	n.b = math.Pow(math.E, a) * math.Sin(b)
	return n
}
