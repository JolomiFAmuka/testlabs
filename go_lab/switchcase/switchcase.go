package main

import (
	"fmt"
)

func main() {

	addsum := func(x, y int) int {
		return x + y
	}
	fmt.Println(addsum(4, 1))

	i := 2
	switch i {
	case 1:
		fmt.Println("it is ONE")
	case 2:
		fmt.Println("it is TWO")
	case 3:
		fmt.Println("it is THREE")
	default:
		fmt.Println("i don't know the number oo")
	}

	whatsmytype := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an integer")
		case float64:
			fmt.Println("I am a float64")
		case string:
			fmt.Println("I am a string")
		default:
			fmt.Println("I don't know your level ooo")
		}
	}

	whatsmytype("Level")
	whatsmytype(true)
	whatsmytype(18.99)
	whatsmytype(8)

	// i := 2
	// fmt.Print("Write ", i, " as ")
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// }

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's the weekend")
	// default:
	// 	fmt.Println("It's a weekday")
	// }

	// t := time.Now()
	// switch {
	// case t.Hour() < 12:
	// 	fmt.Println("It's before noon")
	// default:
	// 	fmt.Println("It's after noon")
	// }

	// whatAmI := func(i interface{}) {
	// 	switch t := i.(type) {
	// 	case bool:
	// 		fmt.Println("I'm a bool")
	// 	case int:
	// 		fmt.Println("I'm an int")
	// 	default:
	// 		fmt.Printf("Don't know type %T\n", t)
	// 	}
	// }
	// whatAmI(true)
	// whatAmI(1)
	// whatAmI("hey")
	// whatAmI('@')
	// whatAmI('q')
}
