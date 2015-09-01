package main

import "fmt"

func rangeMain() {
	ary := []int{10, 20, 30}

	fmt.Print("\"k, v :=range ary\"  ")
	for k, v := range ary {
		fmt.Printf("(k:%v,v:%v)", k, v)
	}
	fmt.Println()

	fmt.Print("\"k:=range ary\"  ")
	for k := range ary {
		fmt.Printf("(k:%v)", k)
	}
	fmt.Println()

	fmt.Print("\"_,v:=range ary\"  ")
	for _, v := range ary {
		fmt.Printf("(v:%v)", v)
	}
	fmt.Println()

	slice := [5]int{11, 12, 13}
	fmt.Print("\"k, v :=range slice\"  ")
	for k, v := range slice {
		fmt.Printf("(k:%v,v:%v)", k, v)
	}
	fmt.Println()

	fmt.Print("\"k:=range slice\"  ")
	for k := range slice {
		fmt.Printf("(k:%v)", k)
	}
	fmt.Println()

	fmt.Print("\"_,v:=range slice\"  ")
	for _, v := range slice {
		fmt.Printf("(v:%v)", v)
	}
	fmt.Println()

	m := map[int]string{1: "a", 2: "b", 3: "b"}
	fmt.Print("\"k, v :=range m\"  ")
	for k, v := range m {
		fmt.Printf("(k:%v,v:%v)", k, v)
	}
	fmt.Println()

	fmt.Print("\"k:=range m\"  ")
	for k := range m {
		fmt.Printf("(k:%v)", k)
	}
	fmt.Println()

	fmt.Print("\"_,v:=range m\"  ")
	for _, v := range m {
		fmt.Printf("(v:%v)", v)
	}
	fmt.Println()

	str := "You中国Hello"
	fmt.Print("\"k, v :=range str\"  ")
	for k, v := range str {
		fmt.Printf("(k:%v,v:%v)", k, v)
	}
	fmt.Println()

	fmt.Print("\"k:=range str\"  ")
	for k := range str {
		fmt.Printf("(k:%v)", k)
	}
	fmt.Println()

	fmt.Print("\"_,v:=range str\"  ")
	for _, v := range str {
		fmt.Printf("(v:%v)", v)
	}
	fmt.Println()

}
