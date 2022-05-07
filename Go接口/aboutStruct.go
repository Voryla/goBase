package main

import "fmt"

func main() {
	var ba = base{Name: "zwk", Width: 180, Height: 180, Derived: derived{
		Height:       0,
		internalFunc: doDerived,
	}}
	ba.baseDo()
	ba.Derived.internalFunc()
	ba.Derived.derivedDo()
	ba.Derived.Height = 3
	ba.Derived.internalFunc()
	ba.derived1.Height = 4
	ba.derived1.internalFunc()
	ba.derived1.internalFunc = doDerived
}

type base struct {
	Name    string
	Width   int
	Height  int
	Derived derived
	derived1
}

func (b base) baseDo() {
	fmt.Println("base func")
}

type derived struct {
	Height       int
	internalFunc func()
}

type derived1 struct {
	Height       int
	internalFunc func()
}

func (d derived) derivedDo() {
	fmt.Println("derived do")
}

func doDerived() {
	fmt.Println("derived do")

}
