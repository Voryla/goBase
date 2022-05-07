package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//f()
	//f1()
	f2()
	//f3()
}

// 考点0
func f() {
	total, sum := 0, 0
	for i := 1; i <= 10; i++ {
		sum += i
		go func() {
			total += i
		}()
	}
	fmt.Printf("total:%d sum %d", total, sum)
}

// 对考点0的改进,但是还是会产生数据竞争
var mutex sync.Mutex

func f2() {
	total, sum := 0, 0
	for i := 1; i <= 10; i++ {
		sum += i
		go func(i int) {
			mutex.Lock()
			total += i
			fmt.Printf("i:%d sum %d\n", total, sum)
			mutex.Unlock()
		}(i)
	}
	fmt.Printf("total:%d sum %d", total, sum)
}

func f3() {
	var wg sync.WaitGroup
	var total int64
	sum := 0
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		sum += i
		go func(i int) {
			defer wg.Done()
			atomic.AddInt64(&total, int64(i))
		}(i)
	}
	wg.Wait()
	fmt.Printf("total:%d sum %d", total, sum)
}
