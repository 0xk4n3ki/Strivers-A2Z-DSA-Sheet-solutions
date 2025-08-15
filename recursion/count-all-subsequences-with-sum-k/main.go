package main
import "fmt"

func main() {
	arr := []int {4, 9, 2, 5, 1}
	k := 10
	cnt, subseq := countSub(arr, k)
	fmt.Printf("arr: %v, count: %d, subseq: %v\n", arr, cnt, subseq)

	ans := printOnlyOne(arr, k)
	fmt.Printf("ans only one: %v\n", ans)
}


// only one subseq
func printOnlyOne(arr []int, target int) []int {
	ans := []int {}
	oneRec(0, arr, target, 0, &ans)

	return ans
}

func oneRec(index int, arr []int, target int, sum int, ans *[]int) bool {
	if target == sum {
		return true
	}else if index >= len(arr) {
		return false
	}

	*ans = append(*ans, arr[index])
	if oneRec(index+1, arr, target, sum+arr[index], ans) {
		return true
	}
	*ans = (*ans)[:len(*ans)-1]

	return  oneRec(index+1, arr, target, sum, ans)
}



// count and all subseq
func countSub(arr []int, target int) (int, [][]int) {
	ans := [][]int {}
	subseq := []int {}
	subRec(0, arr, target, 0, &subseq, &ans)

	count := countRec(0, arr, target, 0)

	return count, ans
}

// generates all subseq
func subRec(index int, arr []int, target int, sum int, subseq *[]int, ans *[][]int) {
	if sum == target {
		tmp := []int {}
		tmp = append(tmp, *subseq...)
		*ans = append(*ans, tmp)
		return
	}else if index == len(arr) {
		return
	}

	*subseq = append(*subseq, arr[index])
	subRec(index+1, arr, target, sum+arr[index], subseq, ans)
	*subseq = (*subseq)[:len(*subseq)-1]
	subRec(index+1, arr, target, sum, subseq, ans)
}

// count the subseq
func countRec(index int, arr []int, target, sum int) int {
	if index == len(arr) {
		if sum == target {
			return 1
		}
		return 0
	}

	l := countRec(index+1, arr, target, sum+arr[index])
	r := countRec(index+1, arr, target, sum)

	return l+r
}