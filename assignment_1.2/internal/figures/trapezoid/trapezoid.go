package trapezoid

type trapezoid struct {
	a, b, h float64
}

func New(a, b, h float64) *trapezoid {
	return &trapezoid{a, b, h}
}

func (tr *trapezoid) Area() float64 {
	return 0.5 * (tr.a + tr.b) * tr.h
}
