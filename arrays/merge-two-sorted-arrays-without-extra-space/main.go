package main

import "fmt"

func main() {
	arr1 := []int {1, 4, 8, 10}
	arr2 := []int {2, 3, 9}

	bruteforceApp(arr1, arr2)

	arr1 = []int {1, 4, 8, 10}
	arr2 = []int {2, 3, 9}
	betterApp1(arr1, arr2)

	arr1 = []int {1, 4, 8, 10}
	arr2 = []int {2, 3, 9}
	betterApp2(arr1, arr2)
}


func betterApp2(arr1, arr2 []int) {
	n, m := len(arr1), len(arr2)

	len := n+m
	gap := (len/2) + (len%2)

	for gap > 0 {
		left := 0
		right := left + gap

		for right < len {
			if left < n && right >= n {
				swapIfGreater(arr1, arr2, left, right-n)
			}else if left >= n {
				swapIfGreater(arr2, arr2, left-n, right-n)
			}else {
				swapIfGreater(arr1, arr1, left, right)
			}
			left++
			right++
		}
		if gap == 1 {
			break
		}
		gap = (gap/2) + (gap%2)
	}

	fmt.Printf("betterApp2:\narr1: %v, arr2: %v\n", arr1, arr2)	
}

func swapIfGreater(arr1, arr2 []int, ind1, ind2 int) {
	if arr1[ind1] > arr2[ind2] {
		arr1[ind1], arr2[ind2] = arr2[ind2], arr1[ind1]
	}
}



func betterApp1(arr1, arr2 []int) {
	n := len(arr1)
	m := len(arr2)

	left, right := n-1, 0

	for left >= 0 && right < m {
		if arr1[left] > arr2[right] {
			arr1[left], arr2[right] = arr2[right], arr1[left]
			left--
			right++
		}else {
			break
		}
	}

	fmt.Printf("arr1: %v, arr2: %v\n", arr1, arr2)

	recBubbleSort(arr1, n)
	recBubbleSort(arr2, m)

	fmt.Printf("after sort: \narr1: %v, arr2: %v\n", arr1, arr2)
}



func bruteforceApp(arr1, arr2 []int) {
	n := len(arr1)
	m := len(arr2)
	arr3 := make([]int, n+m)

	c1, c2 := 0, 0
	i := 0

	for c1 < n && c2 < m {
		if arr1[c1] < arr2[c2] {
			arr3[i] = arr1[c1]
			i++
			c1++
		}else {
			arr3[i] = arr2[c2]
			i++
			c2++
		}
	}

	for c1 < n {
		arr3[i] = arr1[c1]
		c1++
		i++
	}
	for c2 < m {
		arr3[i] = arr2[c2]
		c2++
		i++
	}

	for j := 0; j < n+m; j++ {
		if j < n {
			arr1[j] = arr3[j]
		}else{
			arr2[j-n] = arr3[j]
		}
	}

	fmt.Printf("arr3: %v\n", arr3)
	fmt.Printf("arr1: %v, arr2: %v\n", arr1, arr2)
}


func recBubbleSort(arr []int, n int) {

	if n == 0 {
		return
	}

	swapped := false

	for j := 0; j < n-1; j++ {
		if arr[j] > arr[j+1] {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}

	if !swapped {
		return
	}

	recBubbleSort(arr, n-1)
}