package main

import "fmt"

func main() {
	fmt.Println(sum(9, 2, 1, 4))
	fmt.Println(sum(7))
}

func sum(numbers ...int) int {
	var sum int = 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}
