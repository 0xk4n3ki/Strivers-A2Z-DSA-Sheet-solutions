package main

import "fmt"

func main() {
	n := 36
	sqrt(n)
	sqrt(28)
}

func sqrt(n int) {
	low, high := 1, n/2
	ans := -1

	for low <= high {
		mid := (low+high)/2

		if mid*mid <= n {
			ans = mid
			low = mid + 1
		}else {
			high = mid - 1
		}
	}
	fmt.Printf("n: %d, sqrt: %d\n", n, ans)
}