package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func test() *User {
	var u = User{"nana", 123}
	return &u
}

func main() {

	var user *User = test()
	user2 := user
	user3 := user2
	user.Name = "AAA"
	user3.Name = "BBB"

	fmt.Println(user.Name)
	fmt.Println(user.Age)
	fmt.Println(&user)
	fmt.Println(&user2)
	fmt.Println(&user3)

	var a int
	var p *int
	a = 3
	p = &a
	*p = 20
	fmt.Println(*p)
	fmt.Println(a)
}
