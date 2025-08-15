package main

import "fmt"

func main() {
	num1 := 36
	num2 := 12

	fmt.Printf("divisors of %d: %v\n", num1, divisors(num1))
	fmt.Printf("divisors of %d: %v\n", num2, divisors(num2))
}

func divisors(a int) []int {
	res := []int{}
	for i := 1; i * i <= a; i++ {
		if a % i == 0 {
			if i*i != a {
				res = append(res, i, a/i)
			}else{
				res = append(res, i)
			}
		}
	}
	return res
}