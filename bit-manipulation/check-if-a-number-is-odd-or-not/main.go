package main
import "fmt"

func main() {
	fmt.Printf("n: %d, ans: %t\n", 7, check(7))
	fmt.Printf("n: %d, ans: %t\n", 0, check(0))
}

func check(n int) bool {
	return n & 1 != 0
}