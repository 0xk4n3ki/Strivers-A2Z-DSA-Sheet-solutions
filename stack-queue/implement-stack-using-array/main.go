package main
import "fmt"

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
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) Top() int {
	if s.IsEmpty() {
		return 0
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Print() {
	for _, item := range s.items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}

func main() {
	Stack1 := new(Stack)

	Stack1.Push(372)
	Stack1.Push(40)
	Stack1.Print()
	Stack1.Pop()
	Stack1.Print()
	fmt.Println(Stack1.IsEmpty())
	Stack1.Pop()
	fmt.Println(Stack1.IsEmpty())
}