package main

import "fmt"

func main() {
	fmt.Println(sum(5, 5))
	fmt.Println(multy(5, 5))
}

func sum(first, second int) int {
	result := first + second
	return result
}

func multy(first, second int) int {
	result := first * second
	return result
}
