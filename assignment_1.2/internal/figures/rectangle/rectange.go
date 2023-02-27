package rectangle

type rectangle struct {
	a, b float64
}

func New(a, b float64) *rectangle {
	return &rectangle{a, b}
}

func (rec *rectangle) Area() float64 {
	return rec.a * rec.b
}
