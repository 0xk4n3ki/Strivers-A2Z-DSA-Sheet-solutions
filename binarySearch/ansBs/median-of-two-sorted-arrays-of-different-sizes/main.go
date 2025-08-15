package main

import "fmt"

func main() {
	arr1 := []int {2, 4, 6}
	arr2 := []int {1, 3, 5}

	mergeArr(arr1, arr2)
	withoutMerge(arr1, arr2)

	arr1 = []int {2, 4, 6}
	arr2 = []int {1, 3}

	mergeArr(arr1, arr2)
	withoutMerge(arr1, arr2)
}

func withoutMerge(arr1, arr2 []int) {
	n, m := len(arr1), len(arr2)

	num1, num2 := -1, -1
	ind1, ind2 := (n+m)/2 - 1, (n+m)/2

	cnt := 0

	i, j := 0, 0

	for i < n && j < m {
		if arr1[i] < arr2[j] {
			if cnt == ind1 {
				num1 = arr1[i]
			}
			if cnt == ind2 {
				num2 = arr1[i]
			}
			i++
		}else {
			if cnt == ind1 {
				num1 = arr2[j]
			}
			if cnt == ind2 {
				num2 = arr2[j]
			}
			j++
		}
		cnt++
	}

	for i < n {
		if cnt == ind1 {
			num1 = arr1[i]
		}
		if cnt == ind2 {
			num2 = arr2[i]
		}
		i++
		cnt++
	}

	for j < m {
		if cnt == ind1 {
			num1 = arr2[j]
		}
		if cnt == ind2 {
			num2 = arr2[j]
		}
		j++
		cnt++
	}

	var median float64

	if (n+m) % 2 == 1 {
		median = float64(num2)	
	}else {
		median = float64(num1 + num2) / 2.0
	}

	fmt.Printf("[without merge] -> arr1: %v, arr2: %v, median: %f\n", arr1, arr2, median)
}




func mergeArr(arr1, arr2 []int) {
	n, m := len(arr1), len(arr2)

	fmt.Printf("arr1: %v, arr2: %v\nafter: ", arr1, arr2)

	left, right := n-1, 0

	for left >= 0 && right < m {
		if arr1[left] > arr2[right] {
			arr1[left], arr2[right] = arr2[right], arr1[left]
		}
		left--
		right++
	}
	quickSort(arr1, 0, len(arr1)-1)
	quickSort(arr2, 0, len(arr2)-1)
	fmt.Printf("arr1: %v, arr2: %v ", arr1, arr2)

	medianIndex := (n+m)/2 
	fmt.Printf("medianIndex: %d float: %f ", medianIndex, float64(n+m)/2.0)
	var ans float64
	if float64(medianIndex) * 2.0 != float64(n+m) {
		fmt.Printf(" here:) ")
		if medianIndex < n {
			ans = float64(arr1[medianIndex])
		}else {
			ans = float64(arr2[medianIndex-n])
		}
	}else {
		num1 := int(medianIndex)
		num2 := num1 - 1
		fmt.Printf(" num1: %d, num2: %d ", num1, num2)
		if num1 < n {
			ans += float64(arr1[num1])
		}else {
			ans += float64(arr2[num1-n])
		}
		if num2 < n {
			ans += float64(arr1[num2])
		}else {
			ans += float64(arr2[num2-n])
		}
		ans /= 2.0
	}
	fmt.Printf("median: %f\n", ans)
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
		for arr[j] > pivot && j >=  low + 1 {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[j], arr[low] = arr[low], arr[j]
	return j
}