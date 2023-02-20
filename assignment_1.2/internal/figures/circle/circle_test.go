package circle_test

import (
	"testing"

	"github.com/Goathy/university-assignments/assignment_1.2/internal/figures/circle"
)

func TestCircle(t *testing.T) {
	tt := []struct {
		desc string
		r    float64
		want float64
	}{
		{
			desc: "Area of 6 should be 18.84955592153876",
			r:    6,
			want: 18.84955592153876,
		},
		{
			desc: "Area of 2.2 should be 6.911503837897546",
			r:    2.2,
			want: 6.911503837897546,
		},
	}
	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			circle := circle.New(tc.r)

			got := circle.Area()

			if tc.want != got {
				t.Errorf("want %v, got %v", tc.want, got)
			}
		})
	}
}
