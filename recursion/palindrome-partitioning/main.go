package main
import "fmt"

func main() {
	s := "aabb"
	ans := partition(s)
	fmt.Printf("s: %s, ans: %v\n", s, ans)
}

func partition(s string) [][]string {
	ans := [][]string {}
	sub := []string {}

	backtracking(0, s, &sub, &ans)
	return ans
}

func backtracking(index int, s string, sub *[]string, ans *[][]string) {
	if index == len(s) {
		*ans = append(*ans, append([]string {}, *sub...))
		return
	}

	for i := index; i < len(s); i++ {
		if checkPalindrome(s[index:i+1]) {
			*sub = append(*sub, s[index:i+1])
			backtracking(i+1, s, sub, ans)
			*sub = (*sub)[:len(*sub)-1]
		}
	}
}

func checkPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}