package main

import "fmt"

type Actor struct {
	Name  string
	HP    int
	Speed float64
}

func newActor(name string, hp int, speed float64) *Actor {
	test := Actor{name, hp, speed}
	return &test
}

func main() {
	var actor = newActor("금토끼", 99, 100)
	testac := actor
	testac.Name = "asdf"
	fmt.Println(actor.Name)
	fmt.Println(testac.Name)

}
