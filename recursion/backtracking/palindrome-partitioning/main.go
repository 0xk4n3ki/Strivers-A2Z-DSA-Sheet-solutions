package main
import "fmt"

func main() {
	s := "aabb"
	res := partition(s)
	fmt.Printf("res: %v\n", res)
}

func partition(s string) [][]string {
	ans := [][]string {}
	path := []string {}

	solve(0, s, &path, &ans)
	return ans
}

func solve(index int, s string, path *[]string, ans *[][]string) {
	if index == len(s) {
		tmp := []string {}
		tmp = append(tmp, *path...)
		*ans = append(*ans, tmp)

		return
	}

	for i := index; i < len(s); i++ {
		if isPalindrome(s, index, i) {
			*path = append(*path, s[index:i+1])
			solve(i+1, s, path, ans)
			*path = (*path)[:len(*path)-1]
		}
	}
}

func isPalindrome(s string, index, i int) bool {
	for index <= i {
		if s[index] != s[i] {
			return false
		}
		index++
		i--
	}
	return true
}