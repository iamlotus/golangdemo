package main

import (
	"fmt"
	"sort"
)

type MyIntSlice []int

// sort.Interface requires Len/Less/Swap methods
func (s MyIntSlice) Len() int { return len(s) }

func (s MyIntSlice) Less(i, j int) bool { return s[i] < s[j] }

func (s MyIntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Stringer
func (s MyIntSlice) String() string {
	str := "{"
	for i, e := range s {
		if i > 0 {
			str += ","
		}
		str += fmt.Sprint(e)
	}

	str += "}"

	return str
}

func sortMain() {

	a := MyIntSlice([]int{3, 1, 4, 2, 5})
	fmt.Println("Before Sort,", a)
	sort.Sort(a)
	fmt.Println("After Sort,", a)

	s := MyIntSlice([]int{3, 1, 4, 2, 5})
	sort.Sort(sort.Reverse(s))
	fmt.Println("After Reverse Sort, s=", s, " int[](s)=", []int(s))

}
