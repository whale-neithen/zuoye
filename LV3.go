package main

import "fmt"

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	num := 11
	if isPrime(num) {
		fmt.Printf("%d是素数。\n", num)
	} else {
		fmt.Printf("%d不是素数。\n", num)
	}
}
