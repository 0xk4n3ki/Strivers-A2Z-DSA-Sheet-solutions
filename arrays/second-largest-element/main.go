package main

import "fmt"

func main() {
	arr1 := []int {1, 2, 3, 7, 7, 5}
	arr2 := []int {1}

	fmt.Println("===========using 2 for loop===========")
	s2, l2 := calc(arr1)
	fmt.Printf("arr: %v, 2nd smallest: %d, 2nd largest: %d\n", arr1, s2, l2)

	s2, l2 = calc(arr2)
	fmt.Printf("arr: %v, 2nd smallest: %d, 2nd largest: %d\n", arr2, s2, l2)

	fmt.Println("\n===========using 1 for loop===========")
	s2, l2 = secondSmallest(arr1, len(arr1)), secondLargest(arr1, len(arr1))
	fmt.Printf("arr: %v, 2nd smallest: %d, 2nd largest: %d\n", arr1, s2, l2)

	s2, l2 = secondSmallest(arr2, len(arr2)), secondLargest(arr2, len(arr2))
	fmt.Printf("arr: %v, 2nd smallest: %d, 2nd largest: %d\n", arr2, s2, l2)
}


// better solution : first min and max, then again find min and max, but not equal to last ones
func calc(arr []int) (int, int) {
	if len(arr) < 2 {
		return -1, -1
	}
	min, max, secMin, secMax := arr[0], arr[0], arr[0], arr[0]

	for i := 0; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
		if max < arr[i] {
			max = arr[i]
		}
	}

	for i := 0; i < len(arr); i++ {
		if secMin > arr[i] && arr[i] != min {
			secMin = arr[i]
		}
		if secMax < arr[i] && arr[i] != max {
			secMax = arr[i]
		}
	}
	return secMin, secMax
}

func secondSmallest(arr []int, n int) int {
	if n < 2 {
		return -1
	}
	min, secMin := arr[0], arr[0]
	for i := 0; i < n; i++ {
		if arr[i] < min {
			secMin = min
			min = arr[i]
		}else if arr[i] < secMin && arr[i] != min {
			secMin = arr[i]
		}
	}
	return secMin
}

func secondLargest(arr []int, n int) int {
	if n < 2 {
		return -1
	}
	max, secMax := arr[0], arr[0]
	for i := 0; i < n; i++ {
		if max < arr[i] {
			secMax = max
			max = arr[i]
		}else if secMax < arr[i] && max != arr[i] {
			secMax = arr[i]
		}
	}
	return secMax
}