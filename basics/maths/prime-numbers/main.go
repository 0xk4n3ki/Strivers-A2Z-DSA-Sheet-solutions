package main

import "fmt"

func main() {
	num1 := 2
	num2 := 25
	num3 := 17

	fmt.Printf("%d is prime: %v\n", num1, check(num1))
	fmt.Printf("%d is prime: %v\n", num2, check(num2))
	fmt.Printf("%d is prime: %v\n", num3, check(num3))
}

func check(a int) bool {
	if a == 2 || a == 3 {
		return true
	}
	for i := 2; i*i <= a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}