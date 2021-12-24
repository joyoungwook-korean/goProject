package main

import (
	"fmt"
)

func test(temp int, rain int) string {
	if temp >= 25 && rain > 80 {
		return "덥고 비가 옵니다"
	} else if temp >= 25 && rain <= 20 {
		return "야외활동 하기 좋습니다"
	} else if temp < 25 && (temp < 10 || rain >= 80) {
		return "야외활동 하기 좋지 않습니다"
	} else {
		return "좋은 날씨 입니다."
	}
}

func main() {
	var temp int
	var rain int
	var re int
	re, _ = fmt.Scanln(&temp, &rain)
	fmt.Println(test(temp, rain))
	fmt.Println(re)
}
