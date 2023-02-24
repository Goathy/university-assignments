package circle

import "math"

type circle struct {
	r float64
}

func New(r float64) *circle {
	return &circle{r}
}

func (c *circle) Area() float64 {
	return math.Pi * c.r * c.r
}
