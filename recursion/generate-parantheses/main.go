package main

import (
	"fmt"
)

func main() {
	n := 3
	ans := genPar(n)
	fmt.Printf("ans: %v\n", ans)
}

func genPar(n int) []string {
	ans := []string {}
	// subset := make([]rune, 0)

	// genRec(0, &subset, &ans)            // sol 1: bruteforce
	// genRec(0, 0, n, &subset, &ans)      // sol 2: using subest, optimised
	genRec(0, 0, n, "", &ans)

	return ans
}


func genRec(open, close, n int, subset string, ans *[]string) {
	if len(subset) == n*2 {
		*ans = append(*ans, subset)
		return
	}

	if open < n {
		genRec(open+1, close, n, subset + "(", ans)
	}
	if close < open {
		genRec(open, close+1, n, subset + ")", ans)
	}
}



// func genRec(open, close, n int, subset *[]rune, ans *[]string) {
// 	if len(*subset) == n*2 {
// 		*ans = append(*ans, string(*subset))
// 		return
// 	}

// 	if open < n {
// 		*subset = append(*subset, '(')
// 		genRec(open+1, close, n, subset, ans)
// 		*subset = (*subset)[:len(*subset)-1]
// 	}

// 	if close < open {
// 		*subset = append(*subset, ')')
// 		genRec(open, close+1, n, subset, ans)
// 		*subset = (*subset)[:len(*subset)-1]
// 	}
// }


// func genRec(index int, subset *[]rune, ans *[]string) {
// 	if index == len(*subset) {
// 		if check(*subset) {
// 			*ans = append(*ans, string(*subset))
// 		}
// 		return
// 	}

// 	(*subset)[index] = '('
// 	genRec(index+1, subset, ans)
// 	(*subset)[index] = ')'
// 	genRec(index+1, subset, ans)
// 	(*subset)[index] = 0
// }

// func check(subset []rune) bool {
// 	cnt := 0

// 	for _, ch := range subset {
// 		if ch == '(' {
// 			cnt++
// 		}else if ch == ')' {
// 			cnt--
// 		}
// 		if cnt < 0 {
// 			return false
// 		}
// 	}

// 	return cnt == 0
// }