package main

import (
	"fmt"
	"math"
)

func main() {
	strs := []string {"flower","flow","flight"}

	fmt.Printf("strs: %v, prefix: %s\n", strs, testFunc(strs))
	fmt.Printf("strs: %v, prefix: %s\n", strs, verticalScanning(strs))
	fmt.Printf("strs: %v, prefix: %s\n", strs, divideAndConq(strs))

	fmt.Printf("strs: %v, prefix: %s\n", strs, binarySearch(strs))
}

func binarySearch(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minLen := minLenEle(strs)

	low, high := 1, minLen

	for low <= high {
		mid := (low + high)/2

		if isCommonPrefix(strs, mid) {
			low = mid + 1
		}else {
			high = mid - 1
		}
	}
	return strs[0][0:high]
}

func isCommonPrefix(strs []string, mid int) bool {
	str1 := strs[0][0:mid]
	for i := 1; i < len(strs); i++ {
		str2 := strs[i]

		for j := 0; j < len(str1); j++ {
			if str1[j] != str2[j] {
				return false
			}
		}
	}
	return true
}


// vertical scanning
func verticalScanning(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]

		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || c != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}


// horizontal scanning
func testFunc(str []string) string {
	if len(str) == 0 {
		return ""
	}

	ans := str[0]

	count := len(ans)

	for i := 1; i < len(str); i++ {
		s := str[i]

		j := 0
		for j < min(len(ans), len(s)) && ans[j] == s[j] {
			j++
		}
		count = min(count, j)
	}
	return ans[:count]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}



// divide and conquer algo
func divideAndConq(strs []string) string {
	if len(strs) < 1 {
		return ""
	}

	return dAc(strs, 0, len(strs)-1)
}

func dAc(strs []string, l, r int) string {
	if l == r {
		return strs[l]
	}else {
		mid := (l + r)/2
		lcpLeft := dAc(strs, l, mid)
		lcpRight := dAc(strs, mid+1, r)

		return LCP(lcpLeft, lcpRight)
	}
}

func LCP(str1, str2 string) string {
	minLen := min(len(str1), len(str2))

	for i := 0; i < minLen; i++ {
		if str1[i] != str2[i] {
			return str1[0:i]
		}
	}
	return str1[0:minLen]
}

func minLenEle( strs []string) int {
	minLen := math.MaxInt

	for _, val := range strs {
		if len(val) < minLen {
			minLen = len(val)
		}
	}
	return minLen
}