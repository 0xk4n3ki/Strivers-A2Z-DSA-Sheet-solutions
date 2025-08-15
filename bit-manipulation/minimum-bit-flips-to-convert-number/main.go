package main
import "fmt"

func main() {
	start, goal := 10, 7
	ans := calc(start, goal)
	fmt.Printf("start: %d, goal: %d, ans: %d\n", start, goal, ans)

	start, goal = 3, 4
	ans = calc(start, goal)
	fmt.Printf("start: %d, goal: %d, ans: %d\n", start, goal, ans)
}

func calc(start, goal int) int {
	xor := start ^ goal

	cnt := 0
	for xor > 0 {
		if xor & 1 == 1 {
			cnt++
		}
		xor = xor >> 1
	}
	return cnt
}