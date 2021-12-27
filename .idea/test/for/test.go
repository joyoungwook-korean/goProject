package main

import (
	"fmt"
)

func test(a int) int {
	return a
}

func main() {
	for a := 10; a >= 1; a-- {
		fmt.Printf(" %d", a)
	}

	for i := 2; i < 10; i++ {

		if i == 3 || i == 4 || i == 5 || i == 6 {
			continue
		}
		for j := 1; j < 10; j++ {
			fmt.Println(i, "*", j, "=", i*j)
		}
	}

	for i := 1; i < 10; i += 2 {
		fmt.Println(i, "*", i, i*i)
	}

	for i := 5; i > 0; i-- {
		for j := i; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
