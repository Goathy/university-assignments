package application

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/Goathy/university-assignments/assignment_1.2/internal/figures/circle"
	"github.com/Goathy/university-assignments/assignment_1.2/internal/figures/rectangle"
	"github.com/Goathy/university-assignments/assignment_1.2/internal/figures/trapezoid"
	"github.com/Goathy/university-assignments/assignment_1.2/internal/figures/triangle"
)

const (
	regexMenu          = `^[0-5]$`
	regexPositiveInt   = `^[1-9]+$`
	regexPositiveFloat = `^[0-9](\.|\,)[0-9]+$`
)

var (
	errExit               = errors.New("exit")
	errProvidedWrongValue = errors.New("provided wrong value")
	errUnknownOption      = errors.New("unknown option")
)

var (
	regexMenuCompiled          = regexp.MustCompile(regexMenu)
	regexPositiveIntCompiled   = regexp.MustCompile(regexPositiveInt)
	regexPositiveFloatCompiled = regexp.MustCompile(regexPositiveFloat)
)

type figure interface {
	Area() float64
}

type application struct {
	scanner bufio.Scanner
	stdout  io.Writer
}

func New(stdin io.Reader, stdout io.Writer) *application {
	scanner := bufio.NewScanner(stdin)
	return &application{*scanner, stdout}
}

func (app *application) Run() error {
	app.menu()

	app.scanner.Scan()

	option, err := app.validateMenu(app.scanner.Text())

	if err != nil {
		return err
	}

	fig, err := app.handle(option)

	if errors.Is(err, errExit) {
		return nil
	}

	if err != nil {
		return err
	}

	fmt.Fprintf(app.stdout, "Result: %.2f\n", fig.Area())
	return nil
}

func (app *application) validateMenu(option string) (string, error) {
	if !regexMenuCompiled.MatchString(option) {
		return "", errUnknownOption
	}

	return option, nil
}

func (app *application) handle(option string) (figure, error) {
	switch option {
	case "1":
		params, err := app.handleInput([]string{"a"})

		if err != nil {
			return nil, err
		}

		return rectangle.New(params[0], params[0]), nil
	case "2":
		params, err := app.handleInput([]string{"a", "b"})

		if err != nil {
			return nil, err
		}

		return rectangle.New(params[0], params[1]), nil
	case "3":
		params, err := app.handleInput([]string{"a", "h"})

		if err != nil {
			return nil, err
		}

		return triangle.New(params[0], params[1]), nil

	case "4":
		params, err := app.handleInput([]string{"r"})

		if err != nil {
			return nil, err
		}

		return circle.New(params[0]), nil
	case "5":
		params, err := app.handleInput([]string{"a", "b", "h"})

		if err != nil {
			return nil, err
		}

		return trapezoid.New(params[0], params[1], params[2]), nil
	default:
		return nil, errExit
	}
}

func (app *application) handleInput(opts []string) (params []float64, err error) {

	for _, option := range opts {
		fmt.Fprintf(app.stdout, "Provide %s: ", option)
		app.scanner.Scan()

		switch {
		case regexPositiveIntCompiled.MatchString(app.scanner.Text()):
			num, err := strconv.ParseFloat(app.scanner.Text(), 64)

			if err != nil {
				return nil, err
			}
			params = append(params, num)
		case regexPositiveFloatCompiled.MatchString(app.scanner.Text()):
			formatedFloat := strings.Replace(app.scanner.Text(), ",", ".", 1)
			num, err := strconv.ParseFloat(formatedFloat, 64)
			if err != nil {
				return nil, err
			}

			params = append(params, num)
		default:
			return nil, errProvidedWrongValue
		}
	}
	return
}

func (app *application) menu() {
	greet := ` *******************FIGURE AREA CALCULATOR******************
 1 - Square Area
 2 - Rectange Area
 3 - Triangle Area
 4 - Circle Area
 5 - Trapezoid Area
 0 - Exit

Option: `

	fmt.Fprint(app.stdout, greet)
}
