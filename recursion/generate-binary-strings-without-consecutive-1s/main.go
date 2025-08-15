package main
import "fmt"

func main() {
	n := 3
	ans := generate(n)
	fmt.Printf("ans: %v\n", ans)
}

func generate(n int) []string {
	ans := []string {}
	sub := make([]rune, n)
	
	for i := 0; i < n; i++ {
		sub[i] = '0'
	}

	genRec(0, &sub, &ans)

	return ans
}

func genRec(index int, sub *[]rune, ans *[]string) {
	if index >= len(*sub) {
		*ans = append(*ans, string(*sub))
		return
	}

	genRec(index+1, sub, ans)

	(*sub)[index] = '1'

	genRec(index + 2, sub, ans)

	(*sub)[index] = '0'
}



// func genRec(index, n int, sub *[]rune, ans *[]string) {
// 	if len(*sub) == n {
// 		if check(sub) {
// 			*ans = append(*ans, string(*sub))
// 		}
// 		return
// 	}

// 	*sub = append(*sub, '0')
// 	genRec(index+1, n, sub, ans)
// 	*sub = (*sub)[:len(*sub)-1]
// 	*sub = append(*sub, '1')
// 	genRec(index, n, sub, ans)
// 	*sub = (*sub)[:len(*sub)-1]
// }

// func check(sub *[]rune) bool {
// 	for i := 0; i < len(*sub); i++ {
// 		if i > 0 && i < len(*sub) - 1 {
// 			if (*sub)[i] == '1' && ((*sub)[i - 1] == '1' ||  (*sub)[i + 1] == '1') {
// 				return false
// 			}
// 		}
// 		if i == 0 {
// 			if (*sub)[i] == '1' && (*sub)[i+1] == '1' {
// 				return false
// 			}
// 		}
// 		if i == len(*sub)-1 {
// 			if (*sub)[i] == '1' && (*sub)[i-1] == '1' {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }