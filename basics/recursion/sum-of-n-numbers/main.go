package main

import "fmt"

func main() {
	n1 := 5
	n2 := 6

	fmt.Printf("sum of first %d numbers: %d\n", n1, sumRec(n1))
	fmt.Printf("sum of first %d numbers: %d\n", n2, sumRec(n2))
}

// func sumRec(n, sum int) int {
// 	if n == 0 {
// 		return sum
// 	}
	
// 	return sumRec(n-1, sum+n)
// }


func sumRec(n int) int {
	if n == 0 {
		return n
	}
	return n + sumRec(n-1)
}