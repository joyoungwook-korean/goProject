package main

import "fmt"

type Direction int

const (
	None Direction = iota
	North
	East
	South
	West
)

func GetDirection(direct Direction) string {
	switch direct {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return "None"
	}
}

func GetDirectionToString(angle float64) Direction {
	switch true {
	case angle >= 315 || angle >= 0 && angle < 45:
		return North
	case angle >= 45 && angle < 135:
		return East
	case angle > 135 && angle < 225:
		return South
	case angle >= 225 && angle < 315:
		return West
	default:
		return None
	}
}

func main() {
	fmt.Println(GetDirection(GetDirectionToString(38.3)))
	fmt.Println(GetDirection(GetDirectionToString(235.8)))
	fmt.Println(GetDirection(GetDirectionToString(94.2)))
	fmt.Println(GetDirection(GetDirectionToString(-30)))
}
