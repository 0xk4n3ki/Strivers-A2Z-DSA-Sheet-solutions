package main

import (
	"fmt"
)

func main() {
	arr := []int {3, 2, 1, 5, 3}

	bruteforceApp(arr)
	usingMap(arr)
	usingArr(arr)

	usingMathEquations(arr)
	usingXor(arr)
}


// 
func usingXor(arr []int) {
	n := len(arr)

	xor := 0

	for i := 0; i < n; i++ {
		xor ^= arr[i]
		xor ^= (i+1)
	}


	// instead of this, can use:
	// num := xor & ~(xor-1)     -> this will give the number (1 << bitNum)
	bitNum := 0
	for {
		if (xor & (1 << bitNum)) != 0 {
			break
		}
		bitNum++
	}

	fmt.Printf("bitnum: %d, xor: %d\n", bitNum, xor)

	ones, zeros := 0, 0

	for _, i := range arr {
		if i & (1 << bitNum) != 0 {
			ones = ones ^ i
		}else {
			zeros = zeros ^ i
		}
	}

	for i := 1; i <= n; i++ {
		if i & (1 << bitNum) != 0 {
			ones = ones ^ i
		}else {
			zeros = zeros ^ i
		}
	}

	count := 0
	for _, i := range arr {
		if i == ones {
			count++
		}
	}
	missing, repeated := -1, -1
	if count == 2 {
		missing = zeros
		repeated = ones
	}else {
		missing = ones
		repeated = zeros
	}

	fmt.Printf("arr: %v, missing: %d, repeated: %d\n", arr, missing, repeated)
}


func usingMathEquations(arr []int) {
	n := len(arr)

	sn := n*(n+1)/2
	s2n := n*(n+1)*(2*n+1)/6


	s, s2 := 0, 0

	for _, val := range arr {
		s += val
		s2 += val * val
	}

	xySub := s - sn // x- repeated, y- missing
	xyPlus := (s2 - s2n)/(s - sn)
	repeated := (xySub + xyPlus)/2
	missing := repeated - xySub

	fmt.Printf("arr: %v, missing: %d, repeated: %d\n", arr, missing, repeated)
}


func usingArr(arr []int) {
	n := len(arr)
	hash := make([]int, n+1)
	
	for _, i := range arr {
		hash[i]++
	}

	missing, repeated := -1, -1

	for i := 1; i <= n; i++ {
		if hash[i] == 2 {
			repeated = i
		}else if hash[i] == 0 {
			missing = i
		}

		if missing != -1 && repeated != -1 {
			break
		}
	}

	fmt.Printf("arr: %v, missing: %d, repeated: %d\n", arr, missing, repeated)
}


func usingMap(arr []int) {
	numMap := map[int]int {}
	missing := -1
	repeated := -1

	for _, i := range arr {
		numMap[i]++
	}

	for i := 1; i <= len(arr); i++ {
		if freq, ok := numMap[i]; ok {
			if freq == 2 {
				repeated = i
			}
		}else {
			missing = i
		}
	}
	fmt.Printf("arr: %v, missing: %d, repeated: %d\n", arr, missing, repeated)
}


func bruteforceApp(arr []int) {
	n := len(arr)
	missing := -1
	repeated := -1
	for i := 1; i <= n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if arr[j] == i {
				count++
			}
		}
		if count == 2 {
			repeated = i
		} else if count == 0 {
			missing = i
		}

		if missing != -1 && repeated != -1 {
			break
		}
	}
	fmt.Printf("arr: %v, missing: %d, repeated: %d\n", arr, missing, repeated)
}