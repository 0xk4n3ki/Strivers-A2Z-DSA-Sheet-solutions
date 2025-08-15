package main
import "fmt"

func main() {
	s := "AABABBA"
	ans := brute(s, 1)
	fmt.Printf("ans: %d\n", ans)


	s = "AACBABAAA"
	ans = optimal(s, 2)
	fmt.Printf("ans: %d\n", ans)
}


func optimal(s string, k int) int {
	i, j := 0, 0
	maxLen := 0
	n := len(s)
	hashArr := make([]int, 26)
	maxF := 0

	for j < n {
		hashArr[s[j]-'A']++
		if hashArr[s[j]-'A'] > maxF {
			maxF = hashArr[s[j]-'A']
		}

		if k < (j-i+1) - maxF {
			hashArr[s[i]-'A']--
			// maxF = 0
			// for _, i := range hashArr {
			// 	if i > maxF {
			// 		maxF = i
			// 	}
			// }
			i++
		}

		if k >= (j-i+1) - maxF {
			if (j-i+1) > maxLen {
				maxLen = j-i+1
			}
		}
		j++
	}
	return maxLen
}


func brute(s string, k int) int {
	maxLen := 0
	n := len(s)
	
	for i := 0; i < n; i++ {
		hashArr := make([]int, 26)
		maxF := 0

		for j := i; j < n; j++ {
			hashArr[s[j]-'A']++
			
			if hashArr[s[j]-'A'] > maxF {
				maxF = hashArr[s[j]-'A']
			}

			t := (j-i+1) - maxF
			// fmt.Printf("i: %d, j: %d, t: %d, hashArr: %v\n", i, j, t, hashArr)
			if t > k {
				break
			}
			if j - i + 1 > maxLen {
				maxLen = j - i + 1
			}
		}
	}
	return maxLen
}
