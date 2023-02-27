package triangle_test

import (
	"testing"

	"github.com/Goathy/university-assignments/assignment_1.2/internal/figures/triangle"
)

func TestTriangle(t *testing.T) {
	tt := []struct {
		desc string
		a, h float64
		want float64
	}{
		{
			desc: "6 x 9 should be 27",
			a:    6,
			h:    9,
			want: 27,
		},
		{
			desc: "0.67 x 7 should be 2.345",
			a:    0.67,
			h:    7,
			want: 2.345,
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			tr := triangle.New(tc.a, tc.h)

			got := tr.Area()

			if tc.want != got {
				t.Errorf("want %v, got %v", tc.want, got)
			}
		})
	}
}
