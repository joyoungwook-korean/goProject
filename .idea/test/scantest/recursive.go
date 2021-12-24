package main

import "fmt"

func printNo(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)

	printNo(n - 1)

	fmt.Println("After", n)
}

func main() {
	//printNo(3);
	fmt.Println(F(9))
}

func F(n int) int {
	if n < 2 {
		return n
	}

	return F(n-2) + F(n-1)

}
