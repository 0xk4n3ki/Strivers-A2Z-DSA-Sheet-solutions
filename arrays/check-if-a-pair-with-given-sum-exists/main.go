package main

import "fmt"

func main () {
	arr := []int {2,6,5,8,11}
	n, target := 5, 14
	i, j := getIndex(arr, n, target)
	fmt.Printf("arr[%d]: %d, arr[%d]: %d, arr: %v\n", i, arr[i], j, arr[j], arr)


	arr = []int {2,6,5,8,11}
	n, target = 5, 15
	i, j = usingHashing(arr, n, target)
	if i != -1 && j != -1 {
		fmt.Printf("arr[%d]: %d, arr[%d]: %d, arr: %v\n", i, arr[i], j, arr[j], arr)
	} else {
		fmt.Println("No pair found with the given sum.")
	}


	arr = []int {2,6,5,8,11}
	n, target = 5, 14
	i, j = twoPointer(arr, n, target)
	if i != -1 && j != -1 {
		fmt.Printf("arr[%d]: %d, arr[%d]: %d, arr: %v\n", i, arr[i], j, arr[j], arr)
	} else {
		fmt.Println("No pair found with the given sum.")
	}
}

func getIndex(arr []int, n, target int) (int, int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}else {
				sum := arr[i] + arr[j]
				if sum == target {
					return i, j
				}
			}
		}
	}
	return -1, -1
}

func usingHashing(arr []int, n, target int) (int, int) {
	m := make(map[int]int)

	for i := 0; i < n; i++ {
		moreNeeded := target - arr[i]

		if _, ok := m[moreNeeded]; ok {
			return m[moreNeeded], i
		}
		m[arr[i]] = i
	}
	return -1, -1
}

func twoPointer(arr []int, n, target int) (int, int) {
	quickSort(arr, 0, len(arr)-1)

	left, right := 0, n-1
	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			return left, right
		}else if sum < target {
			left++
		}else {
			right--
		}
	}
	return -1, -1
}


func quickSort(arr []int, low, high int) {
	if low < high {
		pIndex := partition(arr, low, high)
		quickSort(arr, low, pIndex-1)
		quickSort(arr, pIndex+1, high)
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
	arr[low], arr[j] = arr[j], arr[low]
	return j
}