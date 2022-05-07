package main

import (
	"fmt"
)

// https://bitbili.net/golang_spec.html#%E5%AD%97%E7%AC%A6%E4%B8%B2%E5%AD%97%E9%9D%A2%E5%80%BC+string+literals
func main() {
	//interpreted()
	raw()
}

//hello \r world  // output
//hello \r world  // input
//hello			  // output
// 原始字符串,使用``反引号
func raw() {
	str1 := `hello \r world`
	fmt.Println(str1)
	var str2 string
	// 若在输入时输入'\r‘，其后的字符序列将被丢弃
	fmt.Scan(&str2)
	fmt.Println(str2)
}

// output
// \u65e5\u672c\u8a9e
// 日本語
// 解释型字符串 “” 双引号
func interpreted() {
	// 原始字符串字面值，不对字符串进行转义
	str1 := `\u65e5\u672c\u8a9e`
	// 解释型字符串，对字符串进行转义
	str2 := "\u65e5\u672c\u8a9e"
	fmt.Println(str1)
	fmt.Println(str2)
}
