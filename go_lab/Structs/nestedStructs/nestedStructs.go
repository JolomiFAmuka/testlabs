package main

import (
	"fmt"
)
//this code deals with nested structs
type wheel struct {
	Radius   int
	Material string
}

type car struct {
	Make       string
	Model      string
	Height     int
	Width      int
	FrontWheel wheel
	BackWheel  wheel
}

func main() {
	fmt.Println("Hello World")

	mycar := car{
		Make:   "audi",
		Model:  "A8",
		Height: 55,
		Width:  50,
		FrontWheel: wheel{
			Radius:   31,
			Material: "teflon",
		},
		BackWheel: wheel{
			Radius:   33,
			Material: "rubber",
		},
	}

	fmt.Printf("My car %s %s is %d high and %d wide", mycar.Make, mycar.Model, mycar.Height, mycar.Width)
	fmt.Println("")
	fmt.Printf("It's wheel are %d inches in front and %d inches behind ", mycar.FrontWheel.Radius, mycar.BackWheel.Radius)
	fmt.Println("")
	fmt.Printf("The front wheels are made of %s, while the back is %s", mycar.FrontWheel.Material, mycar.BackWheel.Material)
}
