package main

import (
	"fmt"
	"sort"
)

func main() {
	student, cookie := []int {1, 2, 3}, []int {1, 1}
	ans := assign(student, cookie)
	fmt.Println(ans)

	student, cookie = []int {1, 2}, []int {1, 2, 3}
	ans = assign(student, cookie)
	fmt.Println(ans)
}

func assign(student, cookie []int) int {
	sort.Ints(student)
	sort.Ints(cookie)

	i, j := 0, 0
	for i < len(student) && j < len(cookie){
		if cookie[j] >= student[i] {
			i++
		}
		j++
	}
	return i
}