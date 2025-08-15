package main
import "fmt"

func main() {
	fmt.Printf("n: %d, ans: %d\n", 5, count(5))
	fmt.Printf("n: %d, ans: %d\n", 15, count(15))

	fmt.Printf("n: %d, ans: %d\n", 5, optimal(5))
	fmt.Printf("n: %d, ans: %d\n", 15, optimal(15))
}

func count(n int) int {
	cnt := 0
	for n > 0 {
		if n & 1 == 1 {
			cnt++
		}
		n /= 2
	}
	return cnt
}


func optimal(n int) int {
	cnt := 0

	for n > 0 {
		n = n & (n-1)
		cnt++
	}
	return cnt
}