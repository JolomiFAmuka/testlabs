package main

import "fmt"

type car struct {
	Make  string
	Model string
	Year  int
}

func main() {
	mycar := car{
		Make:  "audi",
		Model: "a7",
		Year:  2023,
	}
	fmt.Printf("my %s %s is model %d", mycar.Make, mycar.Model, mycar.Year)

	yourcar := car{}
	yourcar.Make = "Mercedes"
	yourcar.Model = "A63mg"
	yourcar.Year = 2023

	fmt.Println("")
	fmt.Println("********")
	fmt.Printf("Your %s %s is model %d", yourcar.Make, yourcar.Model, yourcar.Year)
}
