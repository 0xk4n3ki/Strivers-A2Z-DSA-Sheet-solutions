package main
import "fmt"

func main() {
	fmt.Printf("n: %d, ans: %d\n", 21, set(21))
	fmt.Printf("n: %d, ans: %d\n", 15, set(15))
}

func set(n int) int {
	if n & (n+1) == 0 {
		return n
	}
	return n | (n+1)
}