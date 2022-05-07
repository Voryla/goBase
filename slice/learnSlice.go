package main

import (
	"fmt"
	"time"
)

func main() {
	//var buffer [256]byte
	//slice := buffer[100:150]
	//fmt.Println("Before: len(slice) =", len(slice))
	//newSlice := SubtractOneFromLength(slice)
	//fmt.Println("After:  len(slice) =", len(slice))
	//fmt.Println("After:  len(newSlice) =", len(newSlice))
	//slice1 := []int{1, 2, 3}
	//slice2 := []int{11, 22, 4, 5}
	//copy(slice1, slice2)
	//fmt.Println(slice1)
	//var m map[string]int
	//m = make(map[string]int)
	//m["one"] = 1
	//fmt.Println(m)
	//var s []int
	//s = make([]int, 0)
	//s = append(s, 1)
	//fmt.Println(s)
	total, sum := 0, 0
	for i := 1; i <= 10; i++ {
		sum += i
		go func() {
			total += i
		}()
	}
	fmt.Printf("total:%d sum %d", total, sum)
}
func SubtractOneFromLength(slice []byte) []byte {
	slice = slice[0 : len(slice)-1]
	return slice
}
func AddOneToEachElement(slice []byte) {
	fmt.Println(len(slice))
	fmt.Printf("%p\n", slice)

	time.Sleep(100 * time.Millisecond)
	fmt.Println(len(slice))
	fmt.Printf("%p\n", slice)

	//fmt.Printf("%p\n", &slice)
	//fmt.Printf("%p\n", slice)

}

func learnSlice() {
	var ints []int
	fmt.Println(ints)
	var floats []float64
	floats = make([]float64, 2, 5)
	fmt.Printf("len: %v, cap: %v\n", len(floats), cap(floats))
	//fmt.Println(floats[2]) // panic 越界
	floats = append(floats, 3)
	fmt.Println(floats[2])

	var array = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := array[1:8]
	fmt.Println(s1[0])
	array[1] = 22
	fmt.Println(s1[0])

	s2 := array[7:]
	fmt.Println(s2)
	// 扩展
	s2 = append(s2, 123)
	fmt.Println(s2)
	array[7] = 777
	fmt.Println(s2)
	fmt.Println(cap(s2))
	fmt.Println(s1)

}

func newSlice() []int {
	arr := [3]int{1, 2, 3}
	slice := arr[0:1]
	return slice
}
