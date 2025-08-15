package main

import (
	"fmt"
	"sort"
)

func main() {
	s1, s2 := "cat", "act"
	afterSort(s1, s2)
	usingFreq(s1, s2)

	s1, s2 = "rules", "lesrt"
	afterSort(s1, s2)
	usingFreq(s1, s2)

	
}

func usingFreq(s1, s2 string) {
	if len(s1) != len(s2) {
		fmt.Printf("s1: %s, s2: %s, ans: %t\n", s1, s2, false)
		return
	}

	freq := [26]int {0}
	for i := 0; i < len(s1); i++ {
		freq[s1[i]-'a']++
	}
	for i := 0; i < len(s1); i++ {
		freq[s2[i]-'a']--
	}

	for i := 0; i < 26; i++ {
		if freq[i] != 0 {
			fmt.Printf("s1: %s, s2: %s, ans: %t\n", s1, s2, false)
			return
		}
	}
	fmt.Printf("s1: %s, s2: %s, ans: %t\n", s1, s2, true)
}


func afterSort(s1, s2 string) {
	if len(s1) != len(s2) {
		fmt.Printf("s1: %s, s2: %s, ans: %t\n", s1, s2, false)
		return
	}

	sl1 := []byte(s1)
	sort.Slice(sl1, func(i, j int) bool {
		return sl1[i] < sl1[j]
	})

	sl2 := []byte(s2)
	sort.Slice(sl2, func(i, j int) bool {
		return sl2[i] < sl2[j]
	})

	if string(sl1) == string(sl2) {
		fmt.Printf("s1: %s, s2: %s, ans: %t\n", s1, s2, true)
		return
	}
	fmt.Printf("s1: %s, s2: %s, ans: %t\n", s1, s2, false)
}