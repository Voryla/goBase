package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	//var b = "Go语言"
	//for i := 0; i < len(b); i++ {
	//	fmt.Printf("%x\n", b[i])
	//}
	//fmt.Println(len(b))
	//// 通过十六进制输出utf8符文
	//fmt.Printf("%c", 0xE0)
	//read()
	//readUtf8()
	//checkString()
	//testStrConv()
	fmt.Println(len(nihongo))
}

const nihongo = "Go语言"

func read() {
	fmt.Printf("%d\n", unsafe.Sizeof(nihongo))
	// 当用range轮询字符串时，轮询的不再是单字节，而是具体的rune。如下所示，对字符串b进行轮询，其第一个参数index代表每个rune的字节偏移量，而
	// runeValue为int32，代表符文数
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}

func readUtf8() {
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U start at byte position %d\n", runeValue, width)
		w = width
	}
}

func checkString() {
	// 判断字符串s是否包含substr字符串
	println(strings.Contains(nihongo, "go"))
	// 判断字符串s是否包含chars字符串中的任一字符
	println(strings.ContainsAny(nihongo, "语abc"))
	// 判断字符串s是否包含符文数r
	println(strings.ContainsRune(nihongo, rune(35328)))
	// 将字符串s以空白字符分割，返回一个切片
	s1 := strings.Fields(nihongo)
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%s\n", s1[i])
	}
	// 将字符串s以满足f(r)==true 的字符分割，返回一个切片
	s := strings.FieldsFunc("a,b,c", func(r rune) bool {
		if strconv.QuoteRune(r) == "," {
			return true
		}
		return false
	})
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", s[i])
	}

	s2 := strings.Split("a,b,c", ",")
	for i := 0; i < len(s2); i++ {
		fmt.Printf("%s\n", s2[i])
	}

}

func testStrConv() {
	// 字符串转换为十进制整数
	println(strconv.Atoi("a"))
	// 字符串转换为某一进制的整数，例如八进制、十六进制
	println(strconv.ParseInt("a", 10, 32))
	// 整数转换为字符串
	println(strconv.Itoa(35328))
	// 某一进制的整数转换为字符串，例如八进制整数转换为字符串
	println(strconv.FormatInt(0xE0, 8))
}

func joinString() {
	// 编译时拼接
	s := "abc" + "abc"
	fmt.Println(s)
	// 运行时拼接
	var a = "hello"
	str := a + "xxs"
	fmt.Println(str)
}
