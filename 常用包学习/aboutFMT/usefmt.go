package main

import (
	"fmt"
	"log"
)

func main() {
	var x uint64 = 1<<64 - 1
	fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))
	log.Println()
}
