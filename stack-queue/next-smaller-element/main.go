package main
import "fmt"

func main() {
	arr := []int {4, 8, 5, 2, 25}
	ans := smaller(arr)
	fmt.Printf("ans: %v\n", ans)

	arr = []int {10, 9, 8, 7}
	ans = smaller(arr)
	fmt.Printf("ans: %v\n", ans)
}

func smaller(arr []int) []int {
	n := len(arr)
	ans := make([]int, n)

	st := new(Stack)

	for i := n-1; i >= 0; i-- {
		for !st.IsEmpty() && st.Top() >= arr[i] {
			st.Pop()
		}
		if st.IsEmpty() {
			ans[i] = -1
		}else {
			ans[i] = st.Top()
		}
		st.Push(arr[i])
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