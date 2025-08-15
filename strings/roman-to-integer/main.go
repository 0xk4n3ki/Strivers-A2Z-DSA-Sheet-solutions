package main

import (
	"fmt"
)

func main() {
	s := "III"
	convert(s)
	simple(s)

	s = "LVIII"
	convert(s)
	simple(s)

	s = "MCMXCIV"
	convert(s)
	simple(s)
}

func simple(s string) {
	rMap := map[string]int {
		"I" : 1,
		"V" : 5,
		"X" : 10,
		"L" : 50,
		"C" : 100,
		"D" : 500,
		"M" : 1000,
	}

	prev := -1
	total := 0

	for i := len(s)-1; i >= 0; i-- {
		curr := rMap[string(s[i])]

		if curr < prev {
			total -= curr
		}else {
			total += curr
		}
		prev = curr
	}
	fmt.Printf("s: %s, ans: %d\n", s, total)
}


func convert(s string) {
	res := 0

	rMap := map[string]int {
		"I" : 1,
		"V" : 5,
		"X" : 10,
		"L" : 50,
		"C" : 100,
		"D" : 500,
		"M" : 1000,
	}

	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			res += rMap[string(s[i])]
			continue
		}
		if s[i] == 'I' && s[i+1] == 'V' {
			res += 4
			i++
		}else if s[i] == 'I' && s[i+1] == 'X' {
			res += 9
			i++
		}else if s[i] == 'X' && s[i+1] == 'L' {
			res += 40
			i++
		}else if s[i] == 'X' && s[i+1] == 'C' {
			res += 90
			i++
		}else if s[i] == 'C' && s[i+1] == 'D' {
			res += 400
			i++
		}else if s[i] == 'C' && s[i+1] == 'M' {
			res += 900
			i++
		}else {
			res += rMap[string(s[i])]
		}
	}
	fmt.Printf("s: %s, res: %d\n", s, res)
}