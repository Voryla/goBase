package main

import "fmt"

func main() {
	learnString()
}

func learnString() {
	var s1 string
	s1 = "hello, 世界"
	fmt.Println(s1)
	fmt.Printf("%c\n", s1[1])
	//s1[0] = 'c'  error 字符串只读，无法直接修改
	// 将字符串转化为slice，此操作会将原字符串底层数组拷贝一份给slice，而非将slice也指向原底层数组
	bs := ([]byte)(s1)
	fmt.Printf("%c\n", bs)
	// 修改slice指定索引的字符
	bs[0] = 'w'
	fmt.Printf("%c\n", bs)
	// 原字符串不受影响
	fmt.Println(s1)
	bw := ([]byte)("世")
	fmt.Println(bw)
}
