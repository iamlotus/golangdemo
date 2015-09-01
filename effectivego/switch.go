package main

import "fmt"

func switchMain() {
	i := 1

	// expression switch
	switch i {
	case 2:
		fmt.Println("i=2")
	case 3:
		fmt.Println("i=3")
	case 4, 5, 6:
		fmt.Println("i in (4,5,6)")
	default:
		fmt.Println("default")
	}

	// no expression switch
	switch {
	default:
		fmt.Println("default")
	case i == 2:
		fmt.Println("i=2")
	case i == 3:
		fmt.Println("i=3")
	case i == 1:
		fmt.Println("i=1")

	}
}
