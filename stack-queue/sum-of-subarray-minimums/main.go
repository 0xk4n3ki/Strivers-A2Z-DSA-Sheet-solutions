package main
import "fmt"

func main() {
	arr := []int {3, 1, 2, 4}
	ans := brute(arr)
	fmt.Printf("ans: %d\n", ans)

	ans = optimal(arr)
	fmt.Printf("ans: %d\n", ans)
}

func optimal(arr []int) int {
	n := len(arr)
	pse := findPSEE(arr)
	nse := findNSE(arr)

	fmt.Printf("pse: %v\n", pse)
	fmt.Printf("nse: %v\n", nse)

	total := 0
	for i := 0; i < n; i++ {
		left := i - pse[i]
		right := nse[i] - i
		total = (total + right*left*arr[i])% MOD
	}
	return total
}

func findNSE(arr []int) []int {
	n := len(arr)
	nse := make([]int, n)
	st := new(Stack)

	for i := n-1; i >= 0; i-- {
		for !st.IsEmpty() && arr[st.Top()] >= arr[i] {
			st.Pop()
		}
		if st.IsEmpty() {
			nse[i] = n
		}else {
			nse[i] = st.Top()
		}
		st.Push(i)
	}
	return nse
}

func findPSEE(arr []int) []int {
	n := len(arr)
	pse := make([]int, n)
	st := new(Stack)

	for i := 0; i < len(arr); i++ {
		for !st.IsEmpty() && arr[st.Top()] > arr[i] {
			st.Pop()
		}
		if st.IsEmpty() {
			pse[i] = -1
		}else {
			pse[i] = st.Top()
		}
		st.Push(i)
	}
	return pse
}


const MOD int = 1_000_000_007
func brute(arr []int) int {
	sum := 0

	for i:= 0; i < len(arr); i++ {
		min := arr[i]
		for j := i; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
			}
			sum = (sum+min)%MOD
		}
	}
	return sum
}















type Stack struct {
	items []int
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Top() int {
	if s.IsEmpty() {
		return 0
	}
	return s.items[s.Size()-1]
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:s.Size()-1]
}