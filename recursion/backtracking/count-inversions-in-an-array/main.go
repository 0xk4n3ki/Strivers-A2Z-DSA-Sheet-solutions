package main
import "fmt"

func main() {
	arr := []int {5, 3, 2, 4, 1}
	count := countInv(arr)
	fmt.Printf("count: %d\n", count)
}

func countInv(arr []int) int {
	ans := mergeSort(arr, 0, len(arr)-1)

	return ans
}

func mergeSort(arr []int, low, high int) int {
	if low >= high {
		return 0
	}
	cnt := 0
	mid := (low + high)/2
	cnt += mergeSort(arr, low, mid)
	cnt += mergeSort(arr, mid+1, high)
	cnt += merge(arr, low, mid, high)

	return cnt
}

func merge(arr []int, low, mid, high int) int {
	i, j := low, mid+1
	tmp := []int {}

	cnt := 0

	for i <= mid && j <= high {
		if arr[i] <= arr[j] {
			tmp = append(tmp, arr[i])
			i++
		}else {
			tmp = append(tmp, arr[j])
			cnt += mid - i + 1
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

	for i = low; i <= high; i++ {
		arr[i] = tmp[i-low]
	}

	return cnt
}