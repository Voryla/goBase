package main

func main() {

}

//	funcName("arg0", 1)
//	funcName("arg0", 3.5)
//	funcName[string, int]("arg0", 4)
//	a := 3
//	b := &a
//	funcName1(b)
//	//funcName("arg0", "string")
//
//	//ints := map[string]int64{
//	//	"first":  34,
//	//	"second": 12,
//	//}
//	//
//	//// Initialize a map for the float values
//	//floats := map[string]float64{
//	//	"first":  35.98,
//	//	"second": 26.99,
//	//}
//	//
//	////strings := map[string]string{
//	////	"first":  "35.98",
//	////	"second": "26.99",
//	////}
//	//fmt.Println(SumIntsOrFloats(ints))
//	//fmt.Println(SumIntsOrFloats(floats))
//	////print(3)
//	////print(float64(3))
//	//// 可以指定以某种类型参数执行
//	//fmt.Printf("Generic Sums: %v and %v \n",
//	//	SumIntsOrFloats[string, int64](ints),
//	//	SumIntsOrFloats[string, float64](floats))
//	////SumIntsOrFloats[string, string](strings)
//	////useInterfaceTypeArgs("s") // error string does not implement Number
//	//useInterfaceTypeArgs(3)
//}
//
//func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
//	var s V
//	for _, v := range m {
//		s += v
//	}
//	return s
//}
//
//type Number interface {
//	float64 | int
//}
//
//func useInterfaceTypeArgs[i Number](s i) {
//	fmt.Println(s)
//}
//
//func print[i string | int | float64](t i) {
//	switch t.(type) {
//	case string:
//		println("string")
//	case int:
//		println("int")
//	}
//}
//func funcName[type0 comparable, type1 Number](arg0 type0, arg1 type1) {
//	//fmt.Printf("arg0 type: %T value: %v\t", arg0, arg0)
//	//fmt.Printf("arg1 type: %T value: %v\n", arg1, arg1)
//	_ = arg0
//	_ = arg1
//}
//
//func funcName1(arg0 interface{}) {
//	_ = arg0
//}
