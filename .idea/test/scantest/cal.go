package main

import "fmt"

func cal(kor int, math int, eng int) float32 {
	tot := float32(kor+math+eng) / float32(3.00)
	return tot
}

func main() {
	tot := cal(100, 70, 53)
	fmt.Printf("tot : %.5f ", tot)
}
