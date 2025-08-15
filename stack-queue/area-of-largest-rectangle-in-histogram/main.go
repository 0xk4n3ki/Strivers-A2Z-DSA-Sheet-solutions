package main
import "fmt"

func main() {
	arr := []int {2, 1, 5, 6, 2, 3}
	ans := largestRectangleArea(arr)
	fmt.Printf("ans: %d\n", ans)

	arr = []int {1, 1}
	ans = largestRectangleArea(arr)
	fmt.Printf("ans: %d\n", ans)

	arr = []int {3, 2, 10, 11, 5, 10, 6, 3}
	ans = optimal(arr)
	fmt.Printf("ans: %d\n", ans)

	arr = []int {2, 1, 5, 6, 2, 3}
	ans = optimal(arr)
	fmt.Printf("ans: %d\n", ans)

	arr = []int {2, 4}
	ans = optimal(arr)
	fmt.Printf("ans: %d\n", ans)

	arr = []int {1}
	ans = optimal(arr)
	fmt.Printf("ans: %d\n", ans)
}

func optimal(arr []int) int {
	st := new(Stack)
	ans := 0
	n := len(arr)

	for i := 0; i < n; i++ {
		for !st.IsEmpty() && arr[st.Top()] > arr[i] {
			height := arr[st.Top()]
			st.Pop()
			width := i-st.Top()-1
			
			if ans < height*width {
				ans = height*width
			}
		}
		st.Push(i)
	}

	for !st.IsEmpty() {
		height := arr[st.Top()]
		st.Pop()

		width := n-st.Top()-1
		if ans < height*width {
			ans = height*width
		}
	}

	return ans
}

func largestRectangleArea(arr []int) int {
	n := len(arr)
	nse := findNSE(arr)
	pse := findPSE(arr)
	fmt.Printf("nse: %v\n", nse)
	fmt.Printf("pse: %v\n", pse)
	
	ans := 0
	for i := 0; i < n; i++ {
		area := arr[i] * (nse[i]-pse[i]-1)
		if area > ans {
			ans = area
		}
	}
	return ans
}

func findNSE(arr []int) []int {
	n := len(arr)
	ans := make([]int, n)
	st := new(Stack)

	for i := n-1; i >= 0; i-- {
		for !st.IsEmpty() && arr[st.Top()] >= arr[i] {
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

func findPSE(arr []int) []int {
	n := len(arr)
	ans := make([]int, n)
	st := new(Stack)

	for i := 0; i < n; i++ {
		for !st.IsEmpty() && arr[st.Top()] >= arr[i] {
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
		return -1
	}
	return s.items[s.Size()-1]
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:s.Size()-1]
}