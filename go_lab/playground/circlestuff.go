package main

import "fmt"

func main() {
	var query int
	var r float64

	fmt.Print("Enter radius: ")
	fmt.Scanf("%f\n", &r)
	fmt.Printf("Enter \n 1 - Area \n 2 - Circumference\n")
	fmt.Scanf("%d\n", &query)

	if query == 1 {
		fmt.Println("The area is", calcArea(r))
	} else {
		fmt.Println("The circumference is", calcCircumference(r))
	}

}

func calcArea(r float64) float64 {
	return 3.14 * r * r
}
func calcCircumference(r float64) float64 {
	return 2 * 3.14 * r
}
