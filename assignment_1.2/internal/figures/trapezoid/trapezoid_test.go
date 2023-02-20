package trapezoid_test

import (
	"testing"

	"github.com/Goathy/university-assignments/assignment_1.2/internal/figures/trapezoid"
)

func TestTrapezoid(t *testing.T) {
	tt := []struct {
		desc    string
		a, b, h float64
		want    float64
	}{
		{
			desc: "Area for a = 6 b = 3.3 h = 0.008 should be 0.037200000000000004",
			a:    6,
			b:    3.3,
			h:    0.008,
			want: 0.037200000000000004,
		},
		{
			desc: "Area for a = 6.3 b = 0.001 h = 0.5 should be 1.57525",
			a:    6.3,
			b:    0.001,
			h:    0.5,
			want: 1.57525,
		},
	}
	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			tr := trapezoid.New(tc.a, tc.b, tc.h)

			got := tr.Area()

			if tc.want != got {
				t.Errorf("want %v, got %v", tc.want, got)
			}
		})
	}
}
