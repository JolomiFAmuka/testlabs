package main

import "fmt"

func main() {

	peepAges := make(map[string]int)

	peepAges["femi"] = 53
	peepAges["mike"] = 49
	peepAges["dayo"] = 41
	peepAges["mervin"] = 42

	fmt.Println(peepAges)

	for x, y := range peepAges {
		fmt.Println(x, y)
	}

}
