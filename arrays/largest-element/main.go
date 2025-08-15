package main

import "fmt"

func main() {
	arr1 := []int {4, 2, 4, 4, 1}
	arr2 := []int {8,10,5,7,9}

	num1 := largest(arr1)
	fmt.Printf("largest element in %v is %d\n", arr1, num1)

	num2 := largest(arr2)
	fmt.Printf("largest element in %v is %d\n", arr2, num2)

	fmt.Println("============using quick sort===========")
	quickSort(arr1, 0, len(arr1)-1)
	fmt.Printf("largest element in %v is %d\n", arr1, arr1[len(arr1)-1])

	quickSort(arr2, 0, len(arr2)-1)
	fmt.Printf("largest element in %v is %d\n", arr2, arr2[len(arr2)-1])
}

func largest(arr []int) int {
	max := arr[0]
	for _, i := range arr {
		if i > max {
			max = i
		}
	}
	return max
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pIndex := partition(arr, low, high)

		quickSort(arr, low, pIndex-1)
		quickSort(arr, pIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := low
	i, j := low, high

	for i < j {
		for i <= high - 1 && arr[i] <= arr[pivot] {
			i++
		}
		for j >= low + 1 && arr[j] > arr[pivot] {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[pivot], arr[j] = arr[j], arr[pivot]
	return j
}