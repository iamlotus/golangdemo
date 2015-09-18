package main

import "fmt"

func mapMain() {
	var s map[int]int
	_, present := s[1]
	fmt.Printf("var s map[int]int => s=%v, s[1]=%v (s==nil)=%v, present=%v\n", s, s[1], s == nil, present)
	s = make(map[int]int)
	_, present1 := s[1]
	fmt.Printf("s = make(map[int]int) => s=%v, s[1]=%v (s==nil)=%v, present=%v\n", s, s[1], s == nil, present1)

	var s1 []int

	fmt.Printf("var s1 []int => s1=%v, (s1==nil)=%v\n", s1, s1 == nil)
	s1 = make([]int, 0)
	fmt.Printf("make([]int,0) => s1=%v, (s1==nil)=%v\n", s1, s1 == nil)
}
