package main

import "fmt"

func main() {
	//number := []int{1, 2, 3, 4}
	//// 删除第一个元素
	//number = number[1:]
	//// 删除最后一个元素
	//number = number[:len(number)-1]
	//
	//// 优雅删除中间一个元素
	//index := len(number) / 2
	//// slice... 为语法糖
	//number = append(number[0:index], number[index+1:]...)
	// 切片复制，两个切片的指针指向同一块数组内存
	old := make([]int64, 3, 3)
	news := old[1:3]
	fmt.Printf("%p\n%p", old, news)
	// 完全复制
	numbers1 := make([]int64, len(old), cap(old)*2)
	count := copy(numbers1, old)
	fmt.Println(count)
	// 切片复制到数组中
	//如果在复制时，数组的长度与切片的长度不同，例如
	//copy（arr[：]，slice），则复制的元素为len（arr）与
	//len（slice）的较小值。
	slice := []byte("abcdefgh")
	var arr [4]byte
	copy(arr[:], slice)
	//copy函数在运行时主要调用了memmove函数，用于实现内存的复
	//制。如果采用协程调用的方式go copy（numbers1，numbers）或者加
	//入了race检测，则会转而调用运行时slicestringcopy或slicecopy函
	//数，进行额外的检查。

}
