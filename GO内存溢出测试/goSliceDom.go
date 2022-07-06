package goMonmOut

func s() {
	//// 创建一个拥有三个int元素的slice
	//slice := []int{1, 2, 3}
	//// 创建一个空切片，make: arg1 type{slice||chan||map}, arg2 大小
	//slice1 := make([]int, 0)

	//  创建一个拥有三个int元素的Array
	array := [5]int{1, 2, 3, 4, 5}
	//array2 := [3]int{}
	//// 创建一个int类型的Array，使用[...] 表示该Array的大小有编译器自行推断
	//array1 := [...]int{1, 2, 3}
	//array3 := new([3]int)
	_ = array
}
