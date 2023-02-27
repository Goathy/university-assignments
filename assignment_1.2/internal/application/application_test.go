package application_test

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/Goathy/university-assignments/assignment_1.2/internal/application"
)

func TestApplication(t *testing.T) {
	tt := []struct {
		desc  string
		input bytes.Buffer
		want  string
	}{
		{
			desc:  "should exit if selected 'exit' option",
			input: *bytes.NewBufferString("0"),
			want:  generateWant(""),
		},
		{
			desc:  "should exit without error if choose 'exit' option",
			input: *bytes.NewBufferString("0"),
			want:  generateWant(""),
		},
	}
	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			stdout := bytes.Buffer{}
			app := application.New(&tc.input, &stdout)

			app.Run()

			got := stdout.String()

			if tc.want != got {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestApplicationErrors(t *testing.T) {
	tt := []struct {
		desc  string
		input bytes.Buffer
		err   error
	}{
		{
			desc:  "should exit if provide unknown option",
			input: *bytes.NewBufferString("6"),
			err:   errors.New("unknown option"),
		},
		{
			desc:  "should exit with error if provide negative number",
			input: *bytes.NewBufferString("1\n-1\n2"),
			err:   errors.New("provided wrong value"),
		},
		{
			desc:  "should exit with error if provide wront option",
			input: *bytes.NewBufferString("6"),
			err:   errors.New("unknown option"),
		},
	}
	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			stdout := bytes.Buffer{}
			app := application.New(&tc.input, &stdout)

			err := app.Run()

			if err.Error() != tc.err.Error() {
				t.Errorf("got %v, want %v", err, tc.err)
			}

		})
	}
}

func TestFigures(t *testing.T) {
	tt := []struct {
		desc  string
		input bytes.Buffer
		want  string
	}{
		{
			desc:  "should calculate square area",
			input: *bytes.NewBufferString("1\n4"),
			want:  generateWant("Provide a: Result: 16.00\n"),
		},
		{
			desc:  "should calculate rectangle area",
			input: *bytes.NewBufferString("2\n4\n6"),
			want:  generateWant("Provide a: Provide b: Result: 24.00\n"),
		},
		{
			desc:  "should calculate triangle area",
			input: *bytes.NewBufferString(("3\n2,3\n5")),
			want:  generateWant("Provide a: Provide h: Result: 5.75\n"),
		},
		{
			desc:  "should calculate circle area",
			input: *bytes.NewBufferString("4\n6.78"),
			want:  generateWant("Provide r: Result: 144.41\n"),
		},
		{
			desc:  "should calculate trapezoid area",
			input: *bytes.NewBufferString("5\n4\n3.334\n7"),
			want:  generateWant("Provide a: Provide b: Provide h: Result: 25.67\n"),
		},
	}
	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			stdout := bytes.Buffer{}
			app := application.New(&tc.input, &stdout)

			app.Run()

			got := stdout.String()

			if tc.want != got {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func generateWant(input string) string {
	greet := ` *******************FIGURE AREA CALCULATOR******************
 1 - Square Area
 2 - Rectange Area
 3 - Triangle Area
 4 - Circle Area
 5 - Trapezoid Area
 0 - Exit

Option: `

	return fmt.Sprintf("%s%s", greet, input)
}
