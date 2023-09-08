package main

import "fmt"

func main() {
	fmt.Println(addemup(10, 20, 30, 40, 50, 60))
}

func addemup(args ...int) int {
	sum := 0

	for _, value := range args {
		sum += value
	}
	return sum
}
