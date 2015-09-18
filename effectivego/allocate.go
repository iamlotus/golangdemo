package main

import (
	"fmt"
	"reflect"
	"sync"
)

type stru struct {
	mutex sync.Mutex
	i     int
	str   string
}

func alloateMain() {
	str := new(string)

	t := reflect.TypeOf(str)
	kind := t.Kind()
	elem := t.Elem()

	fmt.Printf("type(str)=%v,kind=%v,elem=%v \n", t, kind, elem)

	//i := 10
	//t = reflect.TypeOf(i)
	//kind = t.Kind()
	//elem = t.Elem()
	//fmt.Printf(" type(i)=%v,kind=%v,elem=%v \n", i, kind, elem)

	s0 := new(stru)
	fmt.Printf("当s0:=new(stru)时\n s0=%v, s0.mutex=%v, s0.i=%v, 0s.str=%v \n", s0, s0.mutex, s0.i, s0.str)

	var s1 stru
	fmt.Printf("当var s1 stru时\n  s1=%v, s1.mutex=%v, s1.i=%v, s1.str=%v \n", s1, s1.mutex, s1.i, s1.str)

	const Enone, Eio, Einval = 1, 3, 4

	fmt.Printf("TypeOf(Enone)=%v,TypeOf(Eio)=%v,TypeOf(Einval)=%v \n", reflect.TypeOf(Enone), reflect.TypeOf(Eio), reflect.TypeOf(Einval))

	a := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	s := []string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}

	fmt.Printf("a=%v, s=%v, m=%v \n", a, s, m)

	var a1 []int
	b1 := make([]int, 10, 100)
	c1 := make([]int, 10, 10)
	var d1 [10]int
	e1 := make([]int, 10)
	fmt.Printf("TypeOf(a1)=%v, a1=%v, len(a1)=%v cap(a1)=%v\n", reflect.TypeOf(a1), a1, len(a1), cap(a1))
	fmt.Printf("TypeOf(b1)=%v, b1=%v, len(b1)=%v cap(b1)=%v\n", reflect.TypeOf(b1), b1, len(b1), cap(b1))
	fmt.Printf("TypeOf(c1)=%v, c1=%v, len(c1)=%v cap(c1)=%v\n", reflect.TypeOf(c1), c1, len(c1), cap(c1))
	fmt.Printf("TypeOf(d1)=%v, d1=%v, len(d1)=%v cap(d1)=%v\n", reflect.TypeOf(d1), d1, len(d1), cap(d1))
	fmt.Printf("TypeOf(e1)=%v, e1=%v, len(e1)=%v cap(e1)=%v\n", reflect.TypeOf(e1), e1, len(e1), cap(e1))

	f1 := [3]int{1, 2, 3}
	g1 := f1
	for i := 0; i < len(f1); i++ {
		f1[i]++
	}

	elem1, elem2, elem3 := 1, 2, 3
	f2 := [3]*int{&elem1, &elem2, &elem3}
	g2 := f2
	for i := 0; i < len(f1); i++ {
		f1[i]++
		*f2[i]++
	}

	fmt.Printf("To verify array is value and copy by value, f1=%v, g1=%v, f2=%v, g2=%v\n", f1, g1, f2, g2)

	ss1 := make([]byte, 5, 10)
	for i, _ := range ss1 {
		ss1[i] = byte(i + 1)
	}

	fmt.Printf("ss1[:]=%v\n", ss1[:])
	fmt.Printf("ss1[:8]=%v\n", ss1[:8])
	fmt.Printf("cap(ss1[:8])=%v,len(ss1[:8])=%v\n", cap(ss1[:8]), len(ss1[:8]))
	fmt.Printf("cap(ss1[2:8])=%v,len(ss1[2:8])=%v\n", cap(ss1[2:8]), len(ss1[2:8]))
	fmt.Printf("ss1[0:0]=%v\n", ss1[0:0])

	//data := []byte{6, 7, 8}

	//fmt.Printf("first round: ss1=%v \n", ss1)
	//ss1 = Append(ss1, data)
	//fmt.Printf("second round: ss1=%v \n", ss1)
	//ss1 = Append(ss1, data)
	//fmt.Printf("third round: ss1=%v \n", ss1)

}

func Append(slice, data []byte) []byte {
	ls, ld := len(slice), len(data)
	if ls+ld > cap(slice) {
		newSlice := make([]byte, (ls+ld)*2)
		copy(newSlice, slice)
		slice = newSlice
	}

	for i, _ := range data {
		fmt.Printf("!! cap(slice)=%v, ls=%v, i=%v ,", cap(slice), ls, i)
		fmt.Printf("slice[0:cap(slice)]=%v,", slice[0:cap(slice)])
		fmt.Printf("slice[ls+i]=%v\n", slice[ls+i])
		//slice[ls+i] = v
	}
	return slice
}
