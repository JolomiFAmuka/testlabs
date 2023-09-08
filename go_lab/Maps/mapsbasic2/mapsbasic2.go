package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	m["k3"] = 22
	m["u2"] = 50

	fmt.Println("map:", m)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	v7 := m["k7"]
	fmt.Println("v7:", v7)

	fmt.Println("len:", len(m))

	delete(m, "u2")
	fmt.Println("map:", m)

	valll, prs := m["k2"]
	fmt.Println("prs:", prs)
	fmt.Println("valll:", valll)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	q := map[string]int{"one": 1, "two": 2}
	fmt.Println("map:", q)
}
