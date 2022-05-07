package main

import (
	"fmt"
	"time"
)

type World struct {
}

func main() {
	officePlace[Boston] = "BeiJing"
	fmt.Printf("hello %s\n", new(World))
	fmt.Printf("hello %s\n", Boston)
	day := time.Now().Weekday()
	fmt.Printf("Hello, %s (%d)\n", day, day)
}

func (w *World) String() string {
	return "世界"
}

type Office int

var officePlace = make(map[Office]string)

const (
	Boston Office = iota
	NewYork
)

func (o Office) String() string {
	return "Google, " + officePlace[o]
}
