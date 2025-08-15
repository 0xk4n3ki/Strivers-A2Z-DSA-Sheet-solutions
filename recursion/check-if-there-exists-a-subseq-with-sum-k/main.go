package main
import "fmt"

func main() {
	arr := []int {1, 2, 3, 4, 5}
	target := 8

	ans := checkSub(arr, target)
	fmt.Printf("arr: %v, target: %d, ans: %t\n", arr, target, ans)

	target = 17
	ans = checkSub(arr, target)
	fmt.Printf("arr: %v, target: %d, ans: %t\n", arr, target, ans)
}

func checkSub(arr []int, target int) bool {
	return checkRec(0, arr, target, 0)
}

func checkRec(index int, arr []int, target, sum int) bool {
	if sum == target && index < len(arr) {
		return true
	}
	if index >= len(arr) {
		return false
	}

	if checkRec(index+1, arr, target, sum+arr[index]) {
		return true
	}

	return checkRec(index+1, arr, target, sum)
}