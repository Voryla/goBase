package main

import (
	"flag"
)

var MaxOutstanding = 3

type Request struct {
	arg []interface{}
}

var sem = make(chan int, MaxOutstanding)

func handle(req *Request) {
	sem <- 1
	doSomeThing(req)
	<-sem
}

func Server(queue chan *Request) {
	for {
		req := <-queue
		go handle(req)
	}
}
func Server1(queue chan *Request) {
	for req := range queue {
		sem <- 1
		go func(req *Request) {
			doSomeThing(req)
			<-sem
		}(req)
	}
}
func Server2(queue chan *Request) {
	for req := range queue {
		req := req
		sem <- 1
		go func() {
			doSomeThing(req)
			<-sem
		}()
	}
}
func doSomeThing(r *Request) {

}
func handle1(queue chan *Request) {
	for r := range queue {
		doSomeThing(r)
	}
}

func Serve(clientRequests chan *Request, quit chan bool) {
	// 启动处理程序
	for i := 0; i < MaxOutstanding; i++ {
		go handle1(clientRequests)
	}
	<-quit // 等待通知退出。
}

func main() {
	flag.String("addr", ":1718", "http service address")  // Q=17, R=18
	flag.String("addr1", ":1718", "http service address") // Q=17, R=18
	flag.Parse()
	flag.Usage()
}
