package main

import "fmt"

func main() {
	n := []int{12, 5, 13, 14, 28, 2, 10}
	fmt.Println(arrcheck(n))
}

func arrcheck(numarr []int) int {
	smallest := numarr[0]
	for i, _ := range numarr {
		if numarr[i] < smallest {
			smallest = numarr[i]
		}
	}
	return smallest
}
