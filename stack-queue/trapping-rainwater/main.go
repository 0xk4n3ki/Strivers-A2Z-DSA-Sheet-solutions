package main
import "fmt"

func main() {
	arr := []int {0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Printf("arr: %v\n", arr)
	ans := rainwater(arr)
	fmt.Printf("ans: %d\n", ans)

	arr = []int {4, 2, 0, 3, 2, 5}
	fmt.Printf("arr: %v\n", arr)
	ans = optimal(arr)
	fmt.Printf("ans: %d\n", ans)
}

func optimal(arr []int) int {
	l, r := 0, len(arr)-1
	leftmax, rightmax := 0, 0
	sum := 0

	for l <= r {
		if arr[l] <= arr[r] {
			if arr[l] >= leftmax {
				leftmax = arr[l]
			}else {
				sum  += leftmax - arr[l]
			}
			l++
		}else {
			if arr[r] >= rightmax {
				rightmax = arr[r]
			}else {
				sum += rightmax - arr[r]
			}
			r--
		}
	}
	return sum
}


func rainwater(arr []int) int {
	n := len(arr)
	prefix, suffix := make([]int, n), make([]int, n)

	max1, max2 := 0, 0
	i, j := 0, n-1
	for i < n && j >= 0 {
		if arr[i] > max1 {
			max1 = arr[i]
		}
		prefix[i] = max1
		if arr[j] > max2 {
			max2 = arr[j]
		}
		suffix[j] = max2
		i++
		j--
	}

	fmt.Printf("prefix: %v\n", prefix)
	fmt.Printf("suffix: %v\n", suffix)

	sum := 0

	for i = 0; i < n; i++ {
		low := -1
		if suffix[i] < prefix[i] {
			low = suffix[i]
		}else {
			low = prefix[i]
		}
		sum += low - arr[i]
	}

	return sum
}