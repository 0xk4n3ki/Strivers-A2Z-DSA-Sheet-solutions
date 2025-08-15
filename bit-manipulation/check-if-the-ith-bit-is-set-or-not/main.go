package main
import "fmt"

func main() {
	n, i := 5, 0
	fmt.Printf("n: %d, i: %d, ans: %t\n", n, i, check(n, i))

	n, i = 10, 1
	fmt.Printf("n: %d, i: %d, ans: %t\n", n, i, check(n, i))
}

func check(n, i int) bool {
	return n & (1 << i) != 0
}