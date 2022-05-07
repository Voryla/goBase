package main

import "fmt"

func main() {
	arr2 := [3]int{}
	fmt.Println(arr2)
	//temp := arr2
	//sli := make([]int, 3)
	//copy(sli, arr2[:])
	//fmt.Printf("temp Addr: %p Value: %v\n", &temp, temp)
	//fmt.Printf("Befor func arr2 Addr: %p Value: %v\n", &arr2, arr2)
	//ChangeArray(&arr2)
	//fmt.Printf("After func arr2 Addr: %p Value: %v\n", &arr2, arr2)
	//arr1 := [5]int{1, 2, 3, 4, 5}

	//arr3 := [3]int{1, 2, 3}
	//arr4 := [3]int{2, 2, 3}
	//if arr3 == arr2 {
	//	fmt.Println("true")
	//} else {
	//	fmt.Println("false")
	//}
	//
	//if arr4 == arr2 {
	//	fmt.Println("true")
	//} else {
	//	fmt.Println("false")
	//}
	//
	//_ = arr1
	//_ = arr2
	//fmt.Println(len(arr1))
	//fmt.Println(cap(arr1))
	//fmt.Println(unsafe.Sizeof(arr1))
	//var arr2 = [4]int{}
	//
	//var arr = [3]int{1, 2, 3}
	//// 变为：
	//var arr [3]int
	//arr[0] = 1
	//arr[1] = 2
	//arr[2] = 3
	//fmt.Println(arr1[0])

}
func ChangeArray(arr *[3]int) {
	arr[0] = 99
	fmt.Printf("In func arr Addr: %p Value: %v\n", &arr, arr)
}
