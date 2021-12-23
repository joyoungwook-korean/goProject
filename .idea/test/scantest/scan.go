package main

import "fmt"

func add(a int, b int) int  {

	c:=a+b;
	fmt.Println(c);
	return c
}

func main()  {
	c:=add(3,4);
	fmt.Println(c);
}
