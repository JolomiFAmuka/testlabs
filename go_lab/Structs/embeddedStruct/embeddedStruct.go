package main

import (
	"fmt"
)

// embedded structs
type car struct {
	Make  string
	Model string
}

type truck struct {
	car
	Suspension string
	offroad    bool
}

func main() {
	landcruiser := truck{
		car: car{
			Make:  "Kia",
			Model: "Sorento",
		},
		Suspension: "Auto",
		offroad:    true,
	}

	fmt.Printf("The vehicle is a %s %s with %s suspension", landcruiser.Make, landcruiser.Model, landcruiser.Suspension)
}
