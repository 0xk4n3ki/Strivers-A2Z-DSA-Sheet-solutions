package main

import "fmt"

func main() {
	check("egg", "add")
	check("foo", "bar")
	check("badc", "baba")

	fmt.Printf("str1: %s, str2: %s, ans: %t\n", "egg", "add", indexCheck("egg", "add"))
	fmt.Printf("str1: %s, str2: %s, ans: %t\n", "foo", "bar", indexCheck("foo", "bar"))
	fmt.Printf("str1: %s, str2: %s, ans: %t\n", "badc", "baba", indexCheck("badc", "baba"))
}


func indexCheck(s1, s2 string) bool {
	index1, index2 := [200]int {}, [200]int {}

	for i := 0; i < len(s1); i++ {
		if index1[s1[i]] != index2[s2[i]] {
			return false
		}
		index1[s1[i]] = i + 1
		index2[s2[i]] = i + 1
	}
	return true
}



func check(str1, str2 string) {
	n1, n2 := len(str1), len(str2)

	if n1 != n2 {
		fmt.Printf("str1: %s, str2: %s, ans: %t\n", str1, str2, false)
		return
	}

	strMap := map[byte]byte {}
	reverseMap := map[byte]byte {}

	for i := 0; i < n1; i++ {
		a, b := str1[i], str2[i]

		if mapped, ok := strMap[a]; ok {
			if mapped != b {
				fmt.Printf("str1: %s, str2: %s, ans: %t\n", str1, str2, false)
				return
			}
		}else {
			if existing, ok := reverseMap[b]; ok && existing != a {
				fmt.Printf("str1: %s, str2: %s, ans: %t\n", str1, str2, false)
				return
			}
			strMap[a] = b
			reverseMap[b] = a
		}
	}
	fmt.Printf("str1: %s, str2: %s, ans: %t\n", str1, str2, true)
}