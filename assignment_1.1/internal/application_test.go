package application_test

import (
	"bytes"
	"testing"

	application "github.com/Goathy/university-assignments/assignment_1.1/internal"
)

type TestCase struct {
	desc  string
	input bytes.Buffer
	want  string
}

const (
	wantUnknown  = "Provide number\n-> Provided input is unknown\n"
	wantZero     = "Provide number\n-> Provided number is zero\n"
	wantPositive = "Provide number\n-> Provided number is positive\n"
	wantNegative = "Provide number\n-> Provided number is negative\n"
)

func TestDecide(t *testing.T) {
	tt := []TestCase{
		/* ZERO INT */
		{
			desc:  "0 should be zero",
			input: *bytes.NewBufferString("0"),
			want:  wantZero,
		},
		{
			desc:  "d0 should be unknown",
			input: *bytes.NewBufferString("d0"),
			want:  wantUnknown,
		},
		{
			desc:  "0c should be unknown",
			input: *bytes.NewBufferString("0c"),
			want:  wantUnknown,
		},
		{
			desc:  "-0 should be zero",
			input: *bytes.NewBufferString("-0"),
			want:  wantZero,
		},
		{
			desc:  "+0 should be zero",
			input: *bytes.NewBufferString("+0"),
			want:  wantZero,
		},
		/* POSITIVE INT */
		{
			desc:  "1 should be positive",
			input: *bytes.NewBufferString("1"),
			want:  wantPositive,
		},
		{
			desc:  "a1 should be unknown",
			input: *bytes.NewBufferString("a1"),
			want:  wantUnknown,
		},
		{
			desc:  "1t should be unkown",
			input: *bytes.NewBufferString("1k"),
			want:  wantUnknown,
		},
		{
			desc:  "34 should be positive",
			input: *bytes.NewBufferString("34"),
			want:  wantPositive,
		},
		{
			desc:  "g34 should be unknown",
			input: *bytes.NewBufferString("g34"),
			want:  wantUnknown,
		},
		{
			desc:  "34m should be unknown",
			input: *bytes.NewBufferString("34m"),
			want:  wantUnknown,
		},
		{
			desc:  "+1 should be positive",
			input: *bytes.NewBufferString("+1"),
			want:  wantPositive,
		},
		{
			desc:  "+34 should be positive",
			input: *bytes.NewBufferString("+34"),
			want:  wantPositive,
		},
		/* NEGATIVE INT */
		{
			desc:  "-1 should be negative",
			input: *bytes.NewBufferString("-1"),
			want:  wantNegative,
		},
		{
			desc:  "v-1 should be unknown",
			input: *bytes.NewBufferString("v-1"),
			want:  wantUnknown,
		},
		{
			desc:  "-j1 should be unknown",
			input: *bytes.NewBufferString("-j1"),
			want:  wantUnknown,
		},
		{
			desc:  "-1n should be unknown",
			input: *bytes.NewBufferString("-1n"),
			want:  wantUnknown,
		},
		{
			desc:  "-67 should be negative",
			input: *bytes.NewBufferString("-67"),
			want:  wantNegative,
		},
		/* ZERO FLOAT */
		{
			desc:  "0.0 should be zero",
			input: *bytes.NewBufferString("0.0"),
			want:  wantZero,
		},
		{
			desc:  "0b0 should be unknown",
			input: *bytes.NewBufferString("0b0"),
			want:  wantUnknown,
		},
		{
			desc:  "0,0 should be zero",
			input: *bytes.NewBufferString("0,0"),
			want:  wantZero,
		},
		{
			desc:  "0.0d",
			input: *bytes.NewBufferString("0.0d"),
			want:  wantUnknown,
		},
		{
			desc:  "u0.0",
			input: *bytes.NewBufferString("u0.0"),
			want:  wantUnknown,
		},
		{
			desc:  "0.00 should be zero",
			input: *bytes.NewBufferString("0.00"),
			want:  wantZero,
		},
		{
			desc:  "-0.0 should be zero",
			input: *bytes.NewBufferString("-0.0"),
			want:  wantZero,
		},
		{
			desc:  "+0.0 should be zero",
			input: *bytes.NewBufferString("+0.0"),
			want:  wantZero,
		},
		/* POSITIVE FLOAT */
		{
			desc:  "0.1 should be positive",
			input: *bytes.NewBufferString("0.1"),
			want:  wantPositive,
		},
		{
			desc:  "+0.1 should be positive",
			input: *bytes.NewBufferString("+0.1"),
			want:  wantPositive,
		},
		{
			desc:  "0,1 should be positive",
			input: *bytes.NewBufferString("0,1"),
			want:  wantPositive,
		},
		{
			desc:  "0.11 should be positive",
			input: *bytes.NewBufferString("0.11"),
			want:  wantPositive,
		},
		{
			desc:  "1.00 should be positive",
			input: *bytes.NewBufferString("1.00"),
			want:  wantPositive,
		},
		{
			desc:  "1.0.0 should be unknown",
			input: *bytes.NewBufferString("1.0.0"),
			want:  wantUnknown,
		},
		{
			desc:  "0k11 should be unknown",
			input: *bytes.NewBufferString("0k11"),
			want:  wantUnknown,
		},
		{
			desc:  "p0,11 should be unknown",
			input: *bytes.NewBufferString("p0,11"),
			want:  wantUnknown,
		},
		{
			desc:  "0,11t should be unknown",
			input: *bytes.NewBufferString("0,11t"),
			want:  wantUnknown,
		},
		/* NEGATIVE FLOAT */
		{
			desc:  "-0.1 should be negative",
			input: *bytes.NewBufferString("-0.1"),
			want:  wantNegative,
		},
		{
			desc:  "-0,1 should be negative",
			input: *bytes.NewBufferString("-0,1"),
			want:  wantNegative,
		},
		{
			desc:  "-0.11 should be negative",
			input: *bytes.NewBufferString("-0.11"),
			want:  wantNegative,
		},
		{
			desc:  "-0,11 should be negative",
			input: *bytes.NewBufferString("-0,11"),
			want:  wantNegative,
		},
		{
			desc:  "-1.00 should be negative",
			input: *bytes.NewBufferString("-1.00"),
			want:  wantNegative,
		},
		{
			desc:  "-1.0.0 should be unknown",
			input: *bytes.NewBufferString("-1.0.0"),
			want:  wantUnknown,
		},
		{
			desc:  "-0k11 should be unknown",
			input: *bytes.NewBufferString("-0k11"),
			want:  wantUnknown,
		},
		{
			desc:  "-p0,11 should be unknown",
			input: *bytes.NewBufferString("-p0,11"),
			want:  wantUnknown,
		},
		{
			desc:  "-0,11t should be unknown",
			input: *bytes.NewBufferString("-0,11t"),
			want:  wantUnknown,
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			output := &bytes.Buffer{}

			app := application.New(&tc.input, output)

			app.Run()

			got := output.String()

			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
