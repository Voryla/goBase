package main

import "fmt"

var (
	int32Value   int32
	nint32Value  int32
	uint32Value  uint32
	float32Value float32
	float64Value float64
	helloworld   string
	mySlice      []byte
	num          [2]int
	boolValue    bool
	trueValue    bool
	falseValue   bool
)

func main() {
	fmt.Println(int32Value)
	fmt.Println(uint32Value)
	fmt.Println(nint32Value)
	fmt.Println(float32Value)
	fmt.Println(float64Value)
	fmt.Println(helloworld)
	fmt.Println(mySlice)
	fmt.Println(num)
	fmt.Println(boolValue)
	fmt.Println(trueValue)
	fmt.Println(falseValue)
}
