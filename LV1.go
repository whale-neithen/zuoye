package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func main() {
	result := add(1, 2)
	fmt.Println(result)
}
