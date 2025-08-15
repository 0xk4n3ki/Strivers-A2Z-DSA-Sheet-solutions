package main
import "fmt"

func main() {
	arr := []int {2, 10, 12, 1, 11}
	ans := nextGreater(arr)
	fmt.Printf("ans: %v\n", ans)
}

func nextGreater(arr []int) []int {
	n := len(arr)
	ans := make([]int, n)
	st := new(Stack)

	for i := 2*n-1; i >= 0; i-- {
		for !st.IsEmpty() && st.Top() <= arr[i%n] {
			st.Pop()
		}
		
		if st.IsEmpty() {
			ans[i%n] = -1
		}else {
			ans[i%n] = st.Top()
		}
		st.Push(arr[i%n])
	}
	return ans
}

type Stack struct {
	items []int
}


func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:s.Size()-1]
}

func (s *Stack) Top() int {
	if s.IsEmpty() {
		return 0
	}
	return s.items[s.Size()-1]
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}