package main
import "fmt"

func main() {
	arr := []int {1, 4, 3, 2}
	ans := brute(arr)
	fmt.Printf("ans: %d\n", ans)

	arr = []int {1, 4, 3, 2}
	ans = optimal(arr)
	fmt.Printf("ans: %d\n", ans)

	arr = []int {1, 3, 3}
	ans = optimal(arr)
	fmt.Printf("ans: %d\n", ans)
}

func optimal(arr []int) int {
	return longest(arr)-smallest(arr)
}

func longest(arr []int) int {
	n := len(arr)
	pge := findPGE(arr)
	nge := findNGE(arr)
	fmt.Printf("pge: %v\n", pge)
	fmt.Printf("nge: %v\n", nge)

	sum := 0
	for i := 0; i < n; i++ {
		sum += ((nge[i]-i)*(i-pge[i])*arr[i])
	}
	return sum
}

func findNGE(arr []int) []int {
	n := len(arr)
	ans := make([]int, n)
	st := new(Stack)

	for i := n-1; i >= 0; i-- {
		for !st.IsEmpty() && arr[st.Top()] <= arr[i] {
			st.Pop()
		}
		if st.IsEmpty() {
			ans[i] = n
		}else {
			ans[i] = st.Top()
		}
		st.Push(i)
	}
	return ans
}

func findPGE(arr []int) []int {
	n := len(arr)
	ans := make([]int, n)
	st := new(Stack)

	for i := 0; i < n; i++ {
		for !st.IsEmpty() && arr[st.Top()] < arr[i] {
			st.Pop()
		}
		if st.IsEmpty() {
			ans[i] = -1
		}else {
			ans[i] = st.Top()
		}
		st.Push(i)
	}
	return ans
}

func smallest(arr []int) int {
	n := len(arr)
	pse := findPSE(arr)
	nse := findNSE(arr)
	fmt.Printf("pse: %v\n", pse)
	fmt.Printf("nse: %v\n", nse)
	sum := 0

	for i := 0; i < n; i++ {
		sum += ((nse[i]-i) * (i-pse[i])*arr[i])
	}
	// fmt.Printf("sum: %d\n", sum)

	return sum
}

func findNSE(arr []int) []int {
	ans := make([]int, len(arr))
	st := new(Stack)

	for i := len(arr)-1; i >= 0; i-- {
		for !st.IsEmpty() && arr[st.Top()] >= arr[i] {
			st.Pop()
		}
		if st.IsEmpty() {
			ans[i] = len(arr)
		}else {
			ans[i] = st.Top()
		}
		st.Push(i)
	}
	return ans
}

func findPSE(arr []int) []int {
	ans := []int {}
	st := new(Stack)

	for i := 0; i < len(arr); i++ {
		for !st.IsEmpty() && arr[st.Top()] > arr[i] {
			st.Pop()
		}
		if st.IsEmpty() {
			ans = append(ans, -1)
		}else {
			ans = append(ans, st.Top())
		}

		st.Push(i)
	}
	return ans
}






func brute(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		max := arr[i]
		min := arr[i]
		for j := i; j < len(arr); j++ {
			if max < arr[j] {
				max = arr[j]
			}
			if min > arr[j] {
				min = arr[j]
			}
			sum += max-min
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