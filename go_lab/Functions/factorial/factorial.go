package main

import "fmt"
// Factorial Function
// example of recursive function
func fac(num int) int {
	if num == 0 {
		return 1 //factorial of 0 always 1
	}
	return num * (fac(num - 1)) // meat of the code
}

func main() {
	fmt.Println(fac(5))
}
