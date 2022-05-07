package main

import (
	"sync"
)

// 因为case处会直接求值，由于我们的case处为：foo <- <-bar，那么首先会对<-bar求值，
// 而当前代码中，并没有想bar channel发送数据，也就导致<-bar一直处于阻塞状态，进而引发死锁。解决办法就是 将case foo <- <-bar:改为
//	case v:= <-bar:
//		foo <- v
func main() {
	var wg sync.WaitGroup
	foo := make(chan int)
	bar := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case foo <- <-bar:
		default:
			println("default")
		}
	}()
	wg.Wait()
}
