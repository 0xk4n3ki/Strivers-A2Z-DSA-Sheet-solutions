package main

import "fmt"

func main() {
	str1 := "abcdcba"
	str2 := "suraj yadav"

	fmt.Printf("%s is palindrome: %v\n", str1, check(str1, 0, len(str1)-1))
	fmt.Printf("%s is palindrome: %v\n", str2, check(str2, 0, len(str2)-1))
}

func check(str string, i, j int) bool {
	if i >= j {
		return true
	}
	if str[i] != str[j] {
		return false
	}
	
	return check(str, i+1, j-1)
}