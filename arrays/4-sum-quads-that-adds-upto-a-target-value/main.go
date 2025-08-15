package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	arr := []int {4,3,1,4,4,2,3,2,3,1}
	k := 9
	bruteForceApp(arr, k)
	betterApp(arr, k)
	optimalApp(arr, k)
}


// first sort
// then fix two pointers
// move tow pointers from start(after first 2) and end
func optimalApp(arr []int, target int) {
	ans := [][]int {}
	strMap := map[string]struct{} {}

	n := len(arr)

	mergeSort(arr, 0, n-1)

	for i := 0; i < n; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		for j := i+1; j < n; j++ {
			if j > i+1 && arr[j] == arr[j-1] {
				continue
			}

			k, l := j+1, n-1
			for k < l {
				sum := arr[i] + arr[j] + arr[k] + arr[l]

				if sum == target {
					tmp := []int {arr[i], arr[j], arr[k], arr[l]}
					key := checkUniqueness(tmp)

					if _, ok := strMap[key]; !ok {
						ans = append(ans, tmp)
						strMap[key] = struct{}{}
					}
					k, l = k + 1, l - 1
					
					for k < l && arr[k] == arr[k-1] {
						k++
					}
					for k < l && arr[l] == arr[l+1] {
						l--
					}
				}else if sum < target {
					k++
				}else {
					l--
				}
			}
		}
	}
	fmt.Printf("ans: %v\n", ans)
}



// 3 pointers
// and maintain a hashset of remaining elements
// and search for the 4th element(remainder) in the hashset
func betterApp(arr []int, target int) {
	n := len(arr)
	ans := [][]int {}
	strMap := map[string]struct{} {}

	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			hashset := map[int]struct{} {}
			for k := j+1; k < n; k++ {

				sum := arr[i] + arr[j] + arr[k]
				last := target - sum

				if _, ok := hashset[last]; ok {
					tmp := []int {arr[i], arr[j], arr[k], last}
					key := checkUniqueness(tmp)

					if _, ok := strMap[key]; !ok {
						ans = append(ans, tmp)
						strMap[key] = struct{}{}
					}
				}
				hashset[arr[k]] = struct{}{}
				// fmt.Printf("i: %d, j: %d, k: %d, hashset: %v, ans: %v\n", i, j, k, hashset, ans)
			}
		}
	}
	fmt.Printf("ans: %v\n", ans)
}


// 4 loops
func bruteForceApp(arr []int, target int) {
	n := len(arr)
	ans := [][]int {}
	strMap := map[string]struct{} {}

	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				for l := k + 1; l < n; l++ {
					sum := arr[i] + arr[j] + arr[k] + arr[l]
					if sum == target {
						tmp := []int {arr[i], arr[j], arr[k], arr[l]}
						key := checkUniqueness(tmp)
						if _, ok := strMap[key]; !ok {
							ans = append(ans, tmp)
							strMap[key] = struct{}{}
						}
					}
				}
			}
		}
	}
	fmt.Printf("ans: %v\n", ans)
}

func checkUniqueness(arr []int) string {
	sort.Ints(arr)
	return strconv.Itoa(arr[0]) + "," + strconv.Itoa(arr[1]) + "," + strconv.Itoa(arr[2]) + "," + strconv.Itoa(arr[3])
}

func mergeSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	mid := (low + high)/2
	mergeSort(arr, low, mid)
	mergeSort(arr, mid+1, high)
	merge(arr, low, mid, high)
}

func merge(arr []int, low, mid, high int) {
	i, j := low, mid+1
	tmp := []int {}

	for i <= mid && j <= high { 
		if arr[i] <= arr[j] {
			tmp = append(tmp, arr[i])
			i++
		}else {
			tmp = append(tmp, arr[j])
			j++
		}
	}
	for i <= mid {
		tmp = append(tmp, arr[i])
		i++
	}
	for j <= high {
		tmp = append(tmp, arr[j])
		j++
	}

	for i := low; i <= high; i++ {
		arr[i] = tmp[i-low]
	}
}