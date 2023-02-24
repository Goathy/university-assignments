package main

import (
	"os"

	"github.com/Goathy/university-assignments/assignment_1.2/internal/application"
)

func main() {
	app := application.New(os.Stdin, os.Stdout)

	app.Run()
}
