package main

import "fmt"

func test_const() {
	const (
		Red int = iota + 1
		Blue
		Greean
	)

	test_constList(Red)
	test_constList(Blue)
	test_constList(Greean)
}

func test_constList(n int) {
	if n == 1 {
		fmt.Println("Red")
	} else if n == 2 {
		fmt.Println("blue")
	} else {
		fmt.Println("green")
	}
}

const Gravity = 9.80665

func main() {

	test_const()
}
