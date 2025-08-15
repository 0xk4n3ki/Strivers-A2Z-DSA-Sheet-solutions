package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	reversell(arr)
	fmt.Printf("arr: %v\n", arr)


	s := "madama"
	t := checkPalindrome(s, 0)
	fmt.Printf("ans: %t\n", t)


	arr = []int {3, 2, 1}
	ans := subSeq(arr)
	fmt.Printf("ans: %v\n", ans)

	arr = []int {1, 2, 1}
	target := 2
	ans = checkSum(arr, target)
	fmt.Printf("ans: %v\n", ans)
}






func checkSum(arr []int, target int) [][]int {
	ans := [][]int {}
	sub := []int {}

	sumRec(arr, target, 0, 0, &sub, &ans)
	return ans
}
// modified code to print only answer out of [[1, 1], [2]]
func sumRec(arr []int, target int, i, sum int, sub *[]int, ans *[][]int) bool {
	if i >= len(arr) {
		if sum == target {
			tmp := make([]int, 0)
			tmp = append(tmp, *sub...)
			*ans = append(*ans, tmp)
			return true
		}
		return false
	}

	*sub = append(*sub, arr[i])
	sum += arr[i]

	if sumRec(arr, target, i+1, sum, sub, ans) {
		return true
	}
	
	*sub = (*sub)[:len(*sub)-1]
	sum -= arr[i]

	return sumRec(arr, target, i+1, sum, sub, ans)
}





// printing subsequences in array
func subSeq(arr []int) [][]int {
	ans := [][]int {}
	sub := []int {}
	subSeqRec(arr, 0, &sub, &ans)
	return ans
}
func subSeqRec(arr []int, i int, sub *[]int, ans *[][]int) {
	if i >= len(arr) {
		tmp := make([]int, 0)
		tmp = append(tmp, *sub...)
		*ans = append(*ans, tmp)
		return
	}
	*sub = append(*sub, arr[i])
	subSeqRec(arr, i+1, sub, ans)

	*sub = (*sub)[:len(*sub)-1]
	subSeqRec(arr, i+1, sub, ans)
}




// checking if a string is palindrome or not
func checkPalindrome(s string, i int) bool{
	if i >= len(s) / 2 {
		return true
	}

	if s[i] != s[len(s)-i-1] {
		return false
	}
	return checkPalindrome(s, i+1)
}




// reversing a array
func reversell(arr []int) {
	reverseRec(arr, 0)
}
func reverseRec(arr []int, i int) {
	if i > len(arr)/2 {
		return
	}
	reverseRec(arr, i+1)
	arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
}