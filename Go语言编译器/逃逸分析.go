package main

import (
	"fmt"
)

func main() {
	//c := f1()
	//*c = 3
	//fmt.Printf("c的地址 %p\n", &c)
	//fmt.Printf("c= %p\n", c)
	//fmt.Printf("f1= %p\n", *c)
	//f2()
	//time.Sleep(time.Second)

	//slice1 := []int{1}
	//f3(slice1)
	//fmt.Println(cap(slice1))
	//fmt.Printf("%p\n", &slice1)

	//slic := make([]int, 0)
	//slic = append(slic, 1)
	//slic = append(slic, 1)
	//fmt.Println(len(slic))
	//demo := 1
	//fmt.Printf("%p", &demo)

	fun := returnFunc()
	fun()
}

func f1() *int {
	i := 3
	fmt.Printf("i= %p\n", &i)
	return &i
}

func f2() {
	q := 3
	h := 2
	go func() {
		fmt.Println(q, h)
	}()
	//q = 99
}

func f3(s []int) {
	fmt.Printf("%p\n", &s)
	s[0] = 99
	s = append(s, 3)
	fmt.Println(cap(s))
	s[1] = 99

}

func returnFunc() func() {
	value := 3
	return func() {
		value++
	}
}
