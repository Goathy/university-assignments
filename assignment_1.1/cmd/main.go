package main

import (
	"os"

	application "github.com/Goathy/university-assignments/assignment_1.1/internal"
)

func main() {
	app := application.New(os.Stdin, os.Stdout)

	app.Run()
}
