package main
import "fmt"

func main() {
	digits := ""
	ans := comb(digits)
	fmt.Printf("digits: %v, ans: %v\n", digits, ans)
}

func comb(digits string) []string {
	ans := []string {}
	data := []string {"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	sub := []rune {}

	if len(digits) == 0 {
		return ans
	}

	backtracking(0, digits, data, &sub, &ans)

	return ans
}

func backtracking(start int, digits string, data []string, sub *[]rune, ans *[]string) {
	if len(*sub) == len(digits) {
		*ans = append(*ans, string(*sub))
		return
	}

	digit := digits[start] - '0'
	for _, c := range data[digit-2] {
		*sub = append(*sub, c)
		backtracking(start+1, digits, data, sub, ans)
		*sub = (*sub)[:len(*sub)-1]
	}
}