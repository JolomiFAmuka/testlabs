package main

import (
	"fmt"

	"example.com/go-demo-1/mascot"
)

func main() {

	fmt.Println(mascot.BestMascot())
	// var slice []int
	// slice = append(slice, 10) // create slice using append function
	// slice = append(slice, 20)
	// slice = append(slice, 30)
	// slice = append(slice, 40)

	// fmt.Println("The slice before reversal is:", slice)
	// for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
	// 	slice[i], slice[j] = slice[j], slice[i] //reverse the slice
	// }

	// fmt.Println("The slice after reversal is:")
	// fmt.Println(slice) // print the reversed slice
}
