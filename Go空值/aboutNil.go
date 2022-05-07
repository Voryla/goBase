package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func (p person) String() string {
	return p.name + string(rune(p.age))
}

// 接口其实存在两个参数，一个是实现改接口的类型type 一个是该type的value
// 故：一个接口的表示形式为  interface(type,value),而仅当interface(nil, nil)时，该接口变量才等于nil
func main() {
	var s1 fmt.Stringer    // Stringer (nil, nil)
	fmt.Println(s1 == nil) // true
	var p *person
	var s fmt.Stringer = p // Stringer (*Person, nil)
	fmt.Println(s == nil)  // false
}

type doError struct {
	msg string
}

func (error doError) Error() string {
	return error.msg
}

// 错误示例，不要返回一个nil指针作为接口的参数因为该接口永不为nil
func do() *doError { // error (*doError, nil)
	var err *doError
	// do something
	// if flag err.msg = "error"
	return err // nil of type *doError
}

func testDo() {
	err := do()             // error (*doError, nil)
	fmt.Println(err == nil) // false
}

// 错误示例，不要返回一个nil指针作为接口的参数因为该接口永不为nil
func doCorrect() *doError { // nil of type *doError
	//var err *doError
	// do something
	// if flag err.msg = "error"
	// return err
	// else
	return nil // 显示的返回nil
}

func testDoCorrect() {
	err := doCorrect()      // nil of type *doError
	fmt.Println(err == nil) // true
}

// 这里包装了返回值，导致返回的error接口又不会nil了
func wrapDo() error { // error (*doError, nil)
	return doCorrect() // nil of type *doError
}

func wrapFalse() {
	err := wrapDo()         // error(*doError, nil)
	fmt.Println(err == nil) // false
}

// 对于一个nil切片，其底层Data = nil len = 0 cap = 0, 对于一个nil切片访问时会触发panic，可以向一个nil切片执行 append 函数
func aboutNilOfSlice() {
	var s []int
	l := len(s) // 0
	c := cap(s) // 0
	_, _ = l, c
	for i, v := range s { // noting
		// zero times
		_, _ = i, v
	}
	_ = s[0] // panic: index out of range
}
