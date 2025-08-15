package main
import "fmt"

func main() {
	fmt.Printf("n: %d, ans: %t\n", 1, check(1))
	fmt.Printf("n: %d, ans: %t\n", 16, check(16))
	fmt.Printf("n: %d, ans: %t\n", 3, check(3))
}

func check(n int) bool {
	return n > 0 && n & (n-1) == 0
}