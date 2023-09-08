package main

import "fmt"

func main() {
	rect1 := rectangle{4, 6}
	rect2 := rectangle{
		length: 33,
		width:  14,
	}

	fmt.Println(rect1.area())
	fmt.Println(rect2.area())
}

type rectangle struct {
	length int
	width  int
}

func (r rectangle) area() int {
	return r.length * r.width
}
