package main

import "fmt"

func main() {
	n, m := 3, 27
	bs(n, m)

	n, m = 4, 69
	bs(n, m)

	n, m = 3, 125
	bs(n, m)
}

func bs(n, m int) {
	ans := -1
	low, high := 0, m/n

	for low <= high {
		mid := (low + high)/2
		
		mul := 1
		overflow := false
		for i := 0; i < n; i++ {
			mul *= mid
			if mul > m {
				overflow = true
				break
			}
		}

		if !overflow && mul == m {
			ans = mid
			break
		}else if overflow {
			high = mid - 1
		}else {
			low = mid + 1
		}
	}
	fmt.Printf("m: %v, %dth root: %d\n", m, n, ans)
}