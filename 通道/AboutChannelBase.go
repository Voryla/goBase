package main

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	//aboutChanSelect()
	//aboutTickChan()
	//testNilChan()
	var x = make([]struct{}, 1000000000)
	fmt.Println(unsafe.Sizeof(x))
}

// 一个良好的合并chan，读取一个closed的chan，会读取到零值，所以需要有效的防止读取零值的情况
func out(out chan<- int, a, b <-chan int) {
	var aClosed, bClosed bool
	for !aClosed || !bClosed {
		select {
		case v, ok := <-a:
			if !ok {
				aClosed = true
				continue
			}
			out <- v
		case v, ok := <-b:
			if !ok {
				bClosed = true
				continue
			}
			out <- v
		}
	}
	close(out)
}

func out1(out chan<- int, a, b <-chan int) {
	for a != nil || b != nil {
		select {
		case v, ok := <-a:
			if !ok {
				a = nil
				continue
			}
			out <- v
		case v, ok := <-b:
			if !ok {
				b = nil
				continue
			}
			out <- v
		}
	}
	close(out)
}

// channel 的三种声明
func aboutStatement() {
	// 可读可写
	var readAndWriteChan chan int
	// 只读
	var onlyReadChan <-chan int
	// 只写
	var onlyWriteChan chan<- int

	readAndWriteChan = make(chan int)
	onlyReadChan = make(chan int)
	onlyWriteChan = make(chan int)
	// 可读可写
	<-readAndWriteChan
	readAndWriteChan <- 1
	// 只读
	//onlyReadChan <- 1
	<-onlyReadChan
	// 只写
	//<-onlyWriteChan
	onlyWriteChan <- 1

	// nil 通道 向未初始化的通道在编译时和运行时并不会报错，不过，显然无法向通道中写入或读取任何数据
	var nilChan chan int
	fmt.Println(nilChan)
}

func aboutChanSelect() {
	//c := make(chan int, 1)
	//c <- 1
	//select {
	//case <-c:
	//	fmt.Println("random 01")
	//case <-c:
	//	fmt.Println("random 02")
	//}

	sl1 := make([]int64, 1023)
	sl2 := make([]int64, 1024*8)
	sl1[0] = 1
	sl2[0] = 1
	fmt.Println(int64(64 * 1024))
}

func aboutTickChan() {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		}
	}
}

// 一个为nil的通道，不管是读取还是写入都将陷入堵塞状态。当select语句的case对nil通道进行操作时，case分支将永远得不到执行。
func testNilChan() {
	a := make(chan int)
	b := make(chan int)
	go func() {
		for i := 0; i < 2; i++ {
			select {
			case a <- 1:
				a = nil
			case b <- 2:
				b = nil
			}
		}
	}()
	fmt.Println(<-a)
	fmt.Println(<-b)
}
