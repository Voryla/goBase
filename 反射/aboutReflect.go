package aboutReflect

import (
	"fmt"
	"reflect"
)

type myInt int

func reflectOfType() {
	var x float64
	t := reflect.TypeOf(x)
	fmt.Println(t)
}

func reflectOfValue() {
	var x float32 = 3.4
	v := reflect.ValueOf(&x).Elem()
	fmt.Println(v)
	fmt.Println(v.String())
	v.Float()
	// 返回接口值
	y := v.Interface().(float32)
	fmt.Println(y)
	fmt.Println(v.CanSet())
	v.SetFloat(3.14)
}

func reflectOfKind() {
	var x = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64：", v.Kind() == reflect.Float64)
	fmt.Println("value", v.Float())
}

func reflectOfKindAndType() {
	var mi myInt = 3
	t := reflect.TypeOf(mi)
	// type 返回静态类型
	fmt.Println(t)
	// kind 返回基本类型
	fmt.Println(t.Kind())
}
