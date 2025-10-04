package main

import "fmt"

func main() {
	s := "abcba"
	ans := substr(s)
	fmt.Printf("ans: %v\n", ans)

	s = "ccabcc"
	ans = substr(s)
	fmt.Printf("ans: %v\n", ans)
}

func substr(s string) []string {
	ans := []string {}
	i, j := 0, 0
	n := len(s)

	charMap := map[byte]int {}

	for j < n {
		charMap[s[j]]++

		for charMap['a'] > 0 && charMap['b'] > 0 && charMap['c'] > 0 {
			tmp := j
			for tmp < n {
				ans = append(ans, s[i:tmp+1])
				tmp++
			}
			charMap[s[i]]--
			i++
		}
		j++
	}


	return ans
}