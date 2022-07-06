package goMonmOut

import "fmt"

type te struct {
	A int16 // 2
	B int8  // 1
	C int64 // 8
}

func SliceExpansion() {
	a := int32(1) //1
	b := int32(2) //4
	c := int32(3) // 4
	fmt.Printf("a addr %p\n", &a)
	fmt.Printf("b addr %p\n", &b)
	fmt.Printf("c addr %p\n", &c)
	s := te{1, 2, 3}
	fmt.Printf("s addr %p\n", &s)
	d := int64(1)
	fmt.Printf("d addr %p\n", &d)
	e := int32(1)
	fmt.Printf("e addr %p\n", &e)
	_, _, _, _ = a, b, c, s
}

func channelBlock() {
	var c chan bool
	c <- true
}
