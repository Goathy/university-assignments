package main

import (
	"fmt"
	"os"

	"github.com/Goathy/university-assignments/assignment_1.2/internal/application"
)

func main() {
	app := application.New(os.Stdin, os.Stdout)

	err := app.Run()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
