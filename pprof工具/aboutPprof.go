package main

import (
	_ "net/http/pprof"
)

func main() {
	//if err := http.ListenAndServe(":6060", nil); err != nil {
	//	log.Fatalln(err)
	//}
	sliInt := make([]int, 1024*8)
	sliInt64 := make([]int64, 1024*8)
	_ = sliInt[0]
	_ = sliInt64[0]
	//sli1 := make([]int8, 1024*8*2)
	//_ = sli1[0]
}
