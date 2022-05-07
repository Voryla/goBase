package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var a string
var c = make(chan int)
var c1 = make(chan int, 1)

func main() {
	// f
	//go f()
	//c <- 0
	//println(a)

	// f1
	//go f1()
	//<-c
	//println(a)

	// f2
	//go f2()
	//<-c
	//println(a)

	//f3
	//go f3()
	//c1 <- 0
	//println(a)

	//f4()
	//// 一定要有一直活动的goroutine,否则会报deadlock
	//select {}

	//l.Lock()
	//go f5()
	//l.Lock()
	//println(str)

	//go f6()
	//time.Sleep(time.Second)
	//go f7()
	//time.Sleep(4 * time.Second)

	//chan1 := make(chan string)
	//go boring("boring!", chan1)
	//
	//chan2 := boring1("hhh")
	//for i := 0; i < 5; i++ {
	//	fmt.Printf("You say: %q\n", <-chan2)
	//}
	//fmt.Println("You're boring; I'm leaving")

	//chan3 := fanIn(boring1("Anna"), boring1("Joe"))
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("You say: %q\n", <-chan3)
	//}
	//
	//chan3 := fanIn1(boring1("Anna"), boring1("Joe"))
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("You say: %q\n", <-chan3)
	//}

	//f8()

	//timer()
	//time.Sleep(2 * time.Second)

	//useQuit()
	//time.Sleep(3 * time.Second)

	//going()

	gogo()
}
func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func boring1(msg string) <-chan string { // Return Receive-Only channel of strings
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// 生成器模式
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

// 生成器模式，与fanIn功能相同，实现不同
func fanIn1(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

// 超时控制模式
func outTime() {
	c := boring1("Joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second): // 若一秒钟过去了，c都没有发送消息，则执行该case,对每条消息进行超时控制
			fmt.Println("time out!")
			return
		}
	}
}

// 定时器模式
func timer() {
	c := boring1("Joe")
	timeout := time.After(1 * time.Second) // 对整个“会话”进行超时控制
	go func() {
		for {
			select {
			case s := <-c:
				fmt.Println(s)
			case <-timeout: // 若一秒钟过去了，c都没有发送消息，则执行该case
				fmt.Println("you're too slow.")
				return
			}
		}
	}()

}

func useQuit() {
	quit := make(chan bool)
	c := boring3("Joe", quit)
	for i := rand.Intn(100); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
	fmt.Println("close")
}

func boring3(msg string, signal chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case c <- fmt.Sprintf("%s: %d", msg, 1):
			case <-signal:
				return
			}
		}
	}()
	return c
}

func going() func() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	fmt.Println("hello")
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	return func() {
		fmt.Println("return func")
	}
}

func f() {
	a = "hello world"
	<-c
}

func f1() {
	a = "hello world"
	c <- 0
}

func f2() {
	a = "hello world"
	close(c)
}

func f3() {
	a = "hello world"
	<-c1
}

var limit = make(chan int, 3)
var work []func()

func f4() {
	work = append(work, doWork)
	work = append(work, doWork)
	work = append(work, doWork)
	work = append(work, doWork)
	work = append(work, doWork)
	work = append(work, doWork)
	work = append(work, doWork)
	work = append(work, doWork)
	work = append(work, doWork)
	for _, w := range work {
		go func() {
			limit <- 1
			w()
			<-limit
		}()
	}
}
func doWork() {
	fmt.Println("do Work！")
}

var l sync.Mutex
var str string

func f5() {
	str = "hello, world"
	l.Unlock()
}

func f6() {
	l.Lock()
	fmt.Println("lock")
	time.Sleep(2 * time.Second)
	fmt.Println("unlock")
	l.Unlock()

}

func f7() {
	l.Lock()
	fmt.Println("f6 get lock")
}

type Message struct {
	msg  string
	wait chan bool
}

func f8() {
	waitForIt := make(chan bool)
	c := make(chan Message)

	go func() {
		for i := 0; i < 5; i++ {
			msg1 := <-c
			fmt.Println(msg1.msg)
			msg2 := <-c
			fmt.Println(msg2.msg)
			msg1.wait <- true
			msg2.wait <- true
		}
	}()

	for i := 0; i < 5; i++ {
		c <- Message{fmt.Sprintf("Anna: %d", i), waitForIt}
		c <- Message{fmt.Sprintf("JO: %d", i), waitForIt}
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		<-waitForIt
		<-waitForIt
	}
}

func f9(c1, c2, c3 chan string) {
	select {
	case v1 := <-c1:
		fmt.Printf("Recevied %v from c1\n", v1)
	case v2 := <-c2:
		fmt.Printf("Recevied %v from c2\n", v2)
	case c3 <- "23":
		fmt.Printf("sent %v to c3\n", 23)
	default:
		fmt.Printf("no one was ready to communicate")

	}
}

// 击鼓传花
func f10(left, right chan int) {
	left <- 1 + <-right
}

func gogo() {
	const n = 100000           // 传递次数
	leftmost := make(chan int) // 创建通道
	right := leftmost          //
	left := leftmost           //
	for i := 0; i < n; i++ {
		right = make(chan int) //
		go f10(left, right)
		left = right
	}
	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}
