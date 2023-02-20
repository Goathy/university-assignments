package rectangle_test

import (
	"testing"

	"github.com/Goathy/university-assignments/assignment_1.2/internal/figures/rectangle"
)

func TestRectangle(t *testing.T) {
	tt := []struct {
		desc string
		a, b float64
		want float64
	}{
		{
			desc: "2 x 3 should be 6",
			a:    2,
			b:    3,
			want: 6,
		},
		{
			desc: "0.1 x 0.2 should be 0.020000000000000004",
			a:    0.1,
			b:    0.2,
			want: 0.020000000000000004,
		},
		{
			desc: "2.0 x 3 should be 6",
			a:    2.0,
			b:    3,
			want: 6,
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {

			rec := rectangle.New(tc.a, tc.b)

			got := rec.Area()
			if tc.want != got {
				t.Errorf("got %v, want %v", got, tc.want)
			}

		})
	}
}
