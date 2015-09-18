package main

import (
	"fmt"
	"io"
)

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

	var s interface{}

	s = int16(1)

	switch s.(type) {
	default:
		fmt.Println("default")
	case int16:
		fmt.Println("s is int16")
	case int32:
		fmt.Println("s is int32")
	case int:
		fmt.Println("s is int")
	}

}

func ReadFull(r io.Reader, buf []byte) (n int, err error) {
	var nn int
	for len(buf) > 0 && err != nil {
		nn, err = r.Read(buf)
		n = n + nn
		buf = buf[nn:]
	}

	return
}
