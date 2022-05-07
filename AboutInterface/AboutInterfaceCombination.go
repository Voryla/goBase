package main

import "fmt"

type Name interface {
	getName() string
}

type Address interface {
	getAddress() string
}

// Information 接口的组合
type Information interface {
	Name
	Address
}

type Person struct {
	name    string
	address string
}

func (p Person) getName() string {
	return p.name
}

func (p Person) getAddress() string {
	return p.address
}

type Car struct {
	name    string
	address string
}

func (c Car) getName() string {
	return c.name
}

func (c Car) getAddress() string {
	return c.address
}

func main() {
	var information Information
	//p := Person{"张三", "河南"}
	//information = p
	information = Person{"张三", "河南"}
	fmt.Println(information.getName())
	fmt.Println(information.getAddress())
	information = Car{"五菱宏光", "河南"}
	fmt.Println(information.getName())
	fmt.Println(information.getAddress())
	//f1()
	f3()
}

// ------------------------------------------------------------------------------------------------------------------
//													类型断言
/*
可以使用语法i.（Type）在运行时获取存储在接口中的类型。其
中i代表接口，Type代表实现此接口的动态类型。下例通过i.
（Person）获取存储在接口中的动态类型结构体。
*/
func f1() {
	// 接口类型断言
	var i Information
	i = Person{"张三", "北京"}
	car := i.(Person)
	fmt.Println(car.address, "\t", car.name)

}

/*
	在编译时会保证类型Type一定是实现了接口i的类型，否则编译不
	会通过。下例试图通过i.（tmp）将接口转换为tmp类型。
*/
func f2() {
	// 接口类型断言
	//var i Information
	//i = Person{"张三", "北京"}
	type temp struct{}
	//car := i.(temp)
	//fmt.Println(car.address, "\t", car.name)

}

/*
	虽然Go语言在编译时已经防止了此类错误，但是仍然需要在运行时判断一次，这是由于在类型断言方法m=i.（Type）中，当Type实现了接口i，
	而接口内部没有任何动态类型（此时为nil）时，在运行时会直接panic，因为nil无法调用任何方法。下例中的接口变量s没有被赋值，
	此时接口内部没有任何动态类型。为了避免运行时报错的尴尬局面，类型转换还有第二种接口类型断言语法。
	i, ok := i.(Person)
*/
func f3() {
	//接口类型断言
	var i Information
	i = Person{"张三", "北京"}
	// i 类型将包含Information类型和Person类型，如果用Information类型变量接收，将返回Information
	i, ok := i.(Person)
	if !ok {
		fmt.Println("i.(Person) is not ok")
	} else {
		fmt.Println(i.getAddress(), "\t", i.getName())
	}
	// i 类型将包含Information类型和Person类型，如果用新变量接收，将返回Person结构体
	car, ok := i.(Person)
	if !ok {
		fmt.Println("i.(Person) is not ok")
	} else {
		fmt.Println(car.address, "\t", car.name)
	}

}

// ------------------------------------------------------------------------------------------------------------------
//
