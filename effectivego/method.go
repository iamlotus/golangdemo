package main

import "fmt"

type Incer interface {
	inc()
}

type Decer interface {
	dec()
}

func invokeIncAndDisplayOfInterface(i Incer) {
	i.inc()
	fmt.Println(i)
}

func invokeDecAndDisplayOfInterface(d Decer) {
	d.dec()
	fmt.Println(d)
}

type Ty int

func (t *Ty) inc() {
	fmt.Printf("[addr(arg)=%p] ", t)
	*t++
}

func (t Ty) dec() {
	fmt.Printf("[addr(arg)=%p] ", &t)
	t--
}

func methodMain() {
	tt := Ty(10)

	// 一系列invoke方法都定义在接口(Incer/Decer)上

	// inc()方法的reciever是指针，不能应用于value
	// 如果用tt作为参数无法通过编译.
	//invokeIncAndDisplayOfInterface(tt)

	invokeIncAndDisplayOfInterface(&tt)

	invokeDecAndDisplayOfInterface(tt)

	// dec()方法的reciever是value，指针也可以
	// 用tt作为参数可以通过编译.
	invokeDecAndDisplayOfInterface(&tt)

	fmt.Printf("tt=%v, addr(tt)=%p\n\n", tt, &tt)
	(&tt).inc()
	fmt.Printf("(&tt).inc(), tt=%v, addr(tt)=%p\n\n", tt, &tt)
	tt.inc()
	fmt.Printf(" tt.inc(), tt=%v, addr(tt)=%p\n\n", tt, &tt)
	(&tt).dec()
	fmt.Printf("(&tt).dec(), tt=%v, addr(tt)=%p\n\n", tt, &tt)
	tt.dec()
	fmt.Printf("tt.dec(), tt=%v, addr(tt)=%p\n\n", tt, &tt)

}
