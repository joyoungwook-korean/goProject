package main

import "fmt"

func main() {
	test := [...]int{12, 23, 44}

	for _, v := range test {
		fmt.Println(v)
	}
}
