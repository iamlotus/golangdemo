package main

import (
	"fmt"
	"reflect"
)

const i = 1 + 3

type ByteSize float64

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fZB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fZB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fZB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fZB", b/KB)
	default:
		return fmt.Sprintf("%.2fB", b)
	}
}

const (
	X  = iota
	KB = ByteSize(1 << (10 * iota))
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
	_ = iota
	R = iota
)

func constMain() {
	fmt.Printf("i=%v, reflect.TypeOf(i)=%v\n", i, reflect.TypeOf(i))
	const i = 1 + 2
	const i64 int64 = 1 + 2
	fmt.Printf("i=%v, reflect.TypeOf(i)=%v\n", i, reflect.TypeOf(i))
	fmt.Printf("i64=%v, reflect.TypeOf(i64)=%v\n", i64, reflect.TypeOf(i64))
	fmt.Printf("X=%v, KB=%v, MB=%v, GB=%v, R=%v\n", X, KB, MB, GB, R)
	fmt.Printf("N=%v, T=%v\n", N, T)
}

const (
	N = 100
	T = iota
)
