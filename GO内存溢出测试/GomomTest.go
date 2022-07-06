package goMonmOut

func test() {
	slice := make([]int, 8192)
	_ = slice
	slice1 := make([]int, 8193)
	_ = slice1
}
