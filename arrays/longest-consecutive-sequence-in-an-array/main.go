package main

import "fmt"

func main() {
	arr := []int {100, 200, 1, 2, 3, 4}
	bruteforceApp(arr)
	betterApp(arr)
	optimalApp(arr)
}

func optimalApp(arr []int) {
	set := make(map[int]struct{})

	for _, val := range arr {
		set[val] = struct{}{}
	}

	longest := 1

	for val := range set {
		if _, ok := set[val-1]; !ok {
			count := 1
			x := val
			for {
				_, ok := set[x+1]
				if !ok {
					break
				}
				x++
				count++
			}
			if longest < count {
				longest = count
			}
		}
	}
	fmt.Printf("length of longest sequence: %d\n", longest)
}

func betterApp(arr []int) {
	quickSort(arr, 0, len(arr)-1)
	fmt.Printf("after sort: %v\n", arr)

	currCount, lastSmallest, longest := 0, -1, 0
	for _, val := range arr {
		if val-1 == lastSmallest {
			currCount++
			lastSmallest = val
		}else if val != lastSmallest {
			currCount = 1
			lastSmallest = val
		}
		if longest < currCount {
			longest = currCount
		}
	}
	fmt.Printf("better approach using sort, longest: %d\n", longest)
}



func bruteforceApp(arr []int) {
	max := 0
	for _, val := range arr {
		count := 1
		val++
		for {
			found := false
			for _, j := range arr {
				if j == val {
					found = true
					count++
					val++
				}
			}
			if !found {
				break
			}
		}
		if count > max {
			max = count
		}
	}
	fmt.Printf("longest consecutive sequence len: %d\n", max)
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
		for i <= high - 1 && arr[i] <= pivot {
			i++
		}
		for j >= low + 1 && arr[j] > pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[j], arr[low] = arr[low], arr[j]
	return j
}