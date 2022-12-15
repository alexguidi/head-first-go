package main

import (
	"fmt"
)

func main() {
	originalCount := 10
	eatenCount := 4

	fmt.Println("I started with", originalCount)
	fmt.Println("Some jerk ate", eatenCount)
	fmt.Println("There are", originalCount-eatenCount)
}
