package main

import (
	"fmt"
)

var count int

func deferMain() {
	fmt.Printf("Before defer count= %v\n", count)
	count = count + 1
	defer fmt.Printf("defer count= %v\n", count)
	fmt.Printf("After defer count= %v\n", count)
	count = count + 1
	defer retrunNil()

}

func retrunNil() interface{} {
	return nil
}
