package main

import (
	"fmt"
	_ "net/http/pprof"
)

type A1 struct {
	W int
}

func main() {
	//m := make(map[string]string)
	//m["z"] = "w"
	//s := m["z"]
	//fmt.Println(s)
	//if err := http.ListenAndServe(":8080", nil); err != nil {
	//	log.Fatalln(err)
	//}
	a := 0xff
	fmt.Printf("%v", a)
	b := -8
	fmt.Printf("%x\n", b)
	c := A1{3}
	fmt.Printf("%%", c)
	d := make([]int, 2)
	d[0] = 1
	fmt.Println(d[0:0])
}
