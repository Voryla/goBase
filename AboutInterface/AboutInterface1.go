package main

import "fmt"

type Shape interface {
	perimeter() float64
	area() float64
}

type Rectangle struct {
	a, b float64
}

func (r Rectangle) perimeter() float64 {
	return (r.a + r.b) * 2
}

func (r Rectangle) area() float64 {
	return r.a * r.b
}
func (r Rectangle) getA() float64 {
	return r.a
}

var myShape Shape

func main() {
	fmt.Println("Shape interface: ", myShape)
	r := Rectangle{3, 4}
	fmt.Println(r.perimeter())
	fmt.Println(r.area())
}
