package triangle

type triangle struct {
	a, h float64
}

func New(a, h float64) *triangle {
	return &triangle{a, h}
}

func (tr *triangle) Area() float64 {
	return 0.5 * tr.a * tr.h
}
