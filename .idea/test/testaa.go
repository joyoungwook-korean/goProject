package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"testgo/testsu"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	randomsu := rand.Intn(100)
	var test2 bool = false
	var cnt int = 0
	for {
		testz, err := inputVal()
		if err == nil {

			str, bool := testsu.Game(testz, randomsu)
			test2 = bool
			println(str)
			cnt++
		}
		if test2 == true {
			println(cnt)
			break
		}
	}
}

var stdin = bufio.NewReader(os.Stdin)

func inputVal() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}
