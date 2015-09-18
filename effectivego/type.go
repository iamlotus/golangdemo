package main

import (
	"fmt"
	"io"
)

type T1 string
type T2 T1
type TT struct {
}

type Struc struct {
}

type header map[string]string

func typeMain() {
	//str := "abc"
	//t1 := T1(str)
	//t2 := T2(t1)
	//fmt.Printf("str=%v,t1=%v, t2=%v\n", str, t1, t2)
	//fmt.Printf("type(str)=%v,type(t1)=%v, type(t2)=%v\n", reflect.TypeOf(str), reflect.TypeOf(t1), reflect.TypeOf(t2))
	//fmt.Printf("type(t1)=%v, type(t2)=%v\n", reflect.TypeOf((t1)), reflect.TypeOf((t2)))

	//var tt TT
	//var tt1 T1
	//var m1 map[int]int
	//m2 := make(map[int]int)
	//fmt.Printf("tt=%v, isNil(tt)=%v\n", tt, isNil(tt))
	//fmt.Printf("tt1=%v, isNil(tt1)=%v\n", tt1, isNil(tt1))
	//fmt.Printf("m1=%v, (m1==nil)=%v,  isNil(m1)=%v\n", m1, m1 == nil, isNil(m1))
	//fmt.Printf("m2=%v, (m2==nil)=%v,  isNil(m2)=%v\n", m2, m2 == nil, isNil(m2))
	//fmt.Printf("nil=%v, isNil(nil)==%v\n", nil, isNil(nil))

	//var e error
	//var e1 error = nil
	//fmt.Printf("e=%v, (e==nil)=%v, isNil(e)=%v\n", e, e == nil, isNil(e))
	//fmt.Printf("e1=%v, (e1==nil)=%v, isNil(e1)=%v, (e==e1)=%v\n", e1, e1 == nil, isNil(e1), e == e1)

	// type
	var string1 string
	// struct
	var strut1 Struc
	// inner type can be nil (slice,map,channel)
	var map1 map[int]int
	map2 := make(map[int]int)
	// pointer
	var ptr *int
	// interface
	var writer io.Writer

	fmt.Printf("string1=%v, isNil(string1)=%v\n", string1, isNil(string1))
	fmt.Printf("strut1=%v, isNil(strut1)=%v\n", strut1, isNil(strut1))
	fmt.Printf("map1=%v, (map1==nil)=%v,  isNil(map1)=%v\n", map1, map1 == nil, isNil(map1))
	fmt.Printf("map2=%v, (map2==nil)=%v,  isNil(map2)=%v\n", map2, map2 == nil, isNil(map2))
	fmt.Printf("ptr=%v, (ptr==nil)=%v,  isNil(ptr)=%v\n", ptr, ptr == nil, isNil(ptr))
	fmt.Printf("writer=%v, (writer==nil)=%v,  isNil(writer)=%v\n", writer, writer == nil, isNil(writer))

	c := code(uint8(8))
	fmt.Printf("Before set c=%#v, address=(%p)\n", c, &c)
	c.set(uint8(9))
	fmt.Printf("After set c=%#v, address=(%p)\n", c, &c)

	h := make(header)
	h["a"] = "b"
	fmt.Printf("h=%v", h)

}

func isNil(i interface{}) bool {
	return i == nil
}

type code uint8

func (c *code) set(val uint8) {
	*c = code(val)
}
