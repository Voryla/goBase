package main

import "fmt"

func main() {
	// 原题
	//var m map[string]int        //A
	//m["a"] = 1
	//if v := m["b"]; v != nil {  //B
	//	fmt.Println(v)
	//}
	m := map[string]int{} //A
	m["a"] = 1
	if v := m["b"]; v != 0 { //B
		fmt.Println(v)
	}
}
