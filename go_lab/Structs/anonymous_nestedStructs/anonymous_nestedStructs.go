package main

import (
	"fmt"
)

// this code deals with ANONYMOUS nested structs
// nested struct is anonymous struct
type car struct {
	Make   string
	Model  string
	Height int
	Width  int
	Wheel  struct {
		Radius   int
		Material string
	}
}

func main() {
	fmt.Println("Hello World")

	mycar := car{
		Make:   "audi",
		Model:  "A8",
		Height: 55,
		Width:  50,
		Wheel: struct {
			Radius   int
			Material string
		}{
			Radius:   33,
			Material: "rubber",
		},
	}

	fmt.Printf("My car %s %s is %d high and %d wide", mycar.Make, mycar.Model, mycar.Height, mycar.Width)
	fmt.Println("")
	fmt.Printf("It's wheel are %d and %s material ", mycar.Wheel.Radius, mycar.Wheel.Material)
	fmt.Println("")
}
