package test99

import "fmt"

const (
	PI = 3.1415
	pi = 3.15151551
)

var ScreenSize int = 1080
var screenHeight int

func PublicFunc() {
	const MyConst = 100
	println("publicFunc", MyConst)
}

type MyInt int

type myString string

type MyStruct struct {
	Age  int
	name string
}

func (m MyStruct) PublicMethod() {
	fmt.Println("this Public method")
}

func (m MyStruct) privateMethod() {
	fmt.Println("this privatemethod")
}
