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
			desc: "Area of 6 should be 113.09733552923255",
			r:    6,
			want: 113.09733552923255,
		},
		{
			desc: "Area of 2.2 should be 15.205308443374602",
			r:    2.2,
			want: 15.205308443374602,
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
