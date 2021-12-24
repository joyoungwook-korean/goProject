package main

import (
	"fmt"
)

func multiRe(a int, b int) (int, bool) {

	if b == 0 {
		return a, false
	} else {
		return a / b, true
	}
}

func selectReturn(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	} else {
		result = a / b
		success = true
		return
	}
}

func main() {
	var b int
	var a int
	n, err := fmt.Scanln(&a, &b)

	if err == nil {
		c, fail := multiRe(a, b)
		fmt.Println(n, c, fail)
	} else {
		fmt.Println(err)
	}

	d, success := multiRe(4, 2)

	fmt.Println(d, success)

	sel_a, sel_success := selectReturn(1, 0)
	sel_b, sel_fail := selectReturn(4, 2)

	fmt.Println(sel_a, sel_success, sel_b, sel_fail)

}
