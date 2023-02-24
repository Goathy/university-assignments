package application

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
)

const (
	regexZero          = `^[\-\+]?0$`
	regexZeroFloat     = `^[\-\+]?(0[\.\,]0+)$`
	regexPositive      = `^[\+]?[1-9]+$`
	regexPositiveFloat = `^[\+]?[0-9](\.|\,)[0-9]+$`
	regexNegative      = `^[\-][1-9]+$`
	regexNegativeFloat = `^[\-][0-9](\.|\,)[0-9]+$`
)

const (
	unknown  = "unknown"
	positive = "positive"
	negative = "negative"
	zero     = "zero"
)

const (
	message_number = "Provided number is"
	message_input  = "Provided input is"
	message_greet  = "Provide number"
	prompt_sign    = "-> "
	eol            = "\n"
)

var (
	regexZeroCompiled          = regexp.MustCompile(regexZero)
	regexZeroFloatCompiled     = regexp.MustCompile(regexZeroFloat)
	regexPositiveCompiled      = regexp.MustCompile(regexPositive)
	regexPositiveFloatCompiled = regexp.MustCompile(regexPositiveFloat)
	regexNegativeCompiled      = regexp.MustCompile(regexNegative)
	regexNegativeFloatCompiled = regexp.MustCompile(regexNegativeFloat)
)

type decision string

type application struct {
	stdin  io.Reader
	stdout io.Writer
}

func New(stdin io.Reader, stdout io.Writer) *application {
	return &application{stdin, stdout}
}

func (app *application) Run() {
	app.greet()
	scanner := bufio.NewScanner(app.stdin)

	scanner.Scan()

	input := scanner.Bytes()

	decision := app.decide(input)

	result := app.prepare(decision)

	app.print(result)
}

func (app *application) decide(input []byte) decision {

	switch {
	case regexZeroCompiled.Match(input) || regexZeroFloatCompiled.Match(input):
		return zero
	case regexPositiveCompiled.Match(input) || regexPositiveFloatCompiled.Match(input):
		return positive
	case regexNegativeCompiled.Match(input) || regexNegativeFloatCompiled.Match(input):
		return negative
	default:
		return unknown
	}
}

func (app *application) greet() {
	fmt.Fprintf(app.stdout, "%s%s", message_greet, eol)
	fmt.Fprint(app.stdout, prompt_sign)
}
func (app *application) prepare(d decision) string {
	switch d {
	case unknown:
		return fmt.Sprintf("%s %s", message_input, d)
	default:
		return fmt.Sprintf("%s %s", message_number, d)
	}
}

func (app *application) print(message string) {
	fmt.Fprintln(app.stdout, message)
}
