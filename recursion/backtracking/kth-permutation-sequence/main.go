package main

import "fmt"

func main() {
	ans := getPerm(4, 17)
	fmt.Printf("ans: %s\n", ans)
}

func getPerm(n, k int) string {
	fact := 1
	numbers := []int{}
	for i := 1; i < n; i++ {
		numbers = append(numbers, i)
		fact *= i
	}
	numbers = append(numbers, n)

	ans := []rune{}
	k = k - 1

	for {
		// fmt.Printf("numbers: %v, k: %d, fact: %d\n", numbers, k, fact)
		ans = append(ans, rune('0'+numbers[k/fact]))
		// tmp := []int {}
		// tmp = append(tmp, numbers[:k/fact]...)
		// tmp = append(tmp, numbers[k/fact+1:]...)
		// numbers = tmp

		numbers = append(numbers[:k/fact], numbers[k/fact+1:]...)

		if len(numbers) == 0 {
			break
		}

		k = k % fact
		fact = fact / len(numbers)
	}
	return string(ans)
}
