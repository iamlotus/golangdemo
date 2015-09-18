package main

import (
	"errors"
	"fmt"
	"reflect"
)

func reflectMain() {
	//var x uint8 = 'x'
	//v := reflect.ValueOf(x)
	//fmt.Println("type:", v.Type())                            // uint8.
	//fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8) // true.
	//x = uint8(v.Uint())

	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("settability of v:", v.CanSet())

	p := reflect.ValueOf(&x)
	fmt.Println("settability of p.element:", p.Elem().CanSet())

	p.Elem().SetFloat(3.5)
	fmt.Println("x=", x)

	var e error = errors.New("ha")
	et := reflect.TypeOf(e)
	fmt.Println("all methods of e:")
	for i := 0; i < et.NumMethod(); i++ {
		m := et.Method(i)
		fmt.Printf("%#v", m)
	}
}
