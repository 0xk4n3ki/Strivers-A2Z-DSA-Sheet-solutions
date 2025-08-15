package main
import "fmt"

func main() {
	fmt.Printf("n: %d, ans: %v\n", 780, calc(780))
	fmt.Printf("n: %d, ans: %v\n", 37, calc(37))

	fmt.Printf("n: %d, ans: %t\n", 37, check(37))
	fmt.Printf("n: %d, ans: %t\n", 36, check(36))
}

func calc(n int) []int {
	ans := []int {}

	for i := 2; i * i <= n; i++ {
		if n % i == 0 {
			ans = append(ans, i)

			for n % i == 0 {
				n = n/i
			}
		}
	}

	if n != 1 {
		ans = append(ans, n)
	}

	return ans
}

func check(n int) bool {
	for i := 2; i * i <= n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}