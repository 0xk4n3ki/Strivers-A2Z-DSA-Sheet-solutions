package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	arr := []int {-1, 0, 1, 2, -1, -4}

	bruteForceApp(arr)
	betterApp(arr)
	optimalApp(arr)
}


func optimalApp(arr []int) {
	quickSort(arr, 0, len(arr)-1)

	ans := [][]int {}

	for i := 0; i < len(arr); i++ {
		if i != 0 && arr[i] == arr[i-1] {
			continue
		}

		j, k := i + 1, len(arr) - 1
		for j < k {
			sum := arr[i] + arr[j] + arr[k]
			if sum < 0 {
				j++
			}else if sum > 0 {
				k--
			}else {
				tmp := []int {arr[i], arr[j], arr[k]}
				ans = append(ans, tmp)
				j++
				k--
				
				for j < k && arr[j] == arr[j-1] {
					j++
				}
				for j < k && arr[k] == arr[k+1] {
					k--
				}
			}
		}
	}
	fmt.Printf("ans: %v\n", ans)
}


func betterApp(arr []int) {
	ans := [][]int {}
	seen := map[string]struct{} {}

	for i := 0; i < len(arr); i++ {
		m := make(map[int]struct{})
		for j := i+1; j < len(arr); j++ {
			third := - (arr[i] + arr[j])

			if _, ok := m[third]; ok {
				tmp := []int {arr[i], arr[j], third}
				key := tripletKey(tmp)

				if _, ok := seen[key]; !ok{
					ans = append(ans, tmp)
					seen[key] = struct{}{}
				}
			}
			m[arr[j]] = struct{}{}
		}
	}
	fmt.Printf("ans: %v\n", ans)
}


func bruteForceApp(arr []int) {
	ans := make([][]int, 0)
	seen := map[string]struct{} {}

	for i := 0; i< len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			for k := j+1; k < len(arr); k++ {
				sum := arr[i] + arr[j] + arr[k]
				
				if sum == 0 {
					t := []int {arr[i], arr[j], arr[k]}
					key := tripletKey(t)

					if _, ok := seen[key]; !ok {
						ans = append(ans, t)
						seen[key] = struct{}{}
					}
				}
			}
		}
	}
	fmt.Printf("ans: %v\n", ans)
}

func tripletKey(t []int) string {
	sort.Ints(t)
	return strconv.Itoa(t[0]) + "," + strconv.Itoa(t[1]) + "," + strconv.Itoa(t[2])
}

func quickSort(arr []int, low, high int) {
	if low < high {
		partitionIndex := partition(arr, low, high)
		quickSort(arr, low, partitionIndex-1)
		quickSort(arr, partitionIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[low]
	i, j := low, high

	for i < j {
		for arr[i] <= pivot && i <= high - 1 {
			i++
		}
		for arr[j] > pivot && j >= low + 1 {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[j], arr[low] = arr[low], arr[j]
	return j
}