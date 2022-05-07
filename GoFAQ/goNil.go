package main

import "fmt"

type err struct {
	a *int
}

func main() {
	var p *err = nil
	if p == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

const (
	Running = iota
	Stopped
	Rebooting
	Terminated
)
