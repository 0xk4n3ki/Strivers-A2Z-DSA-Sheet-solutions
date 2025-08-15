package main
import "fmt"

func main() {
	s := "*+pq-mn"
	ans := toInfix(s)
	fmt.Printf("ans: %s\n", ans)
}

func toInfix(s string) string {
	st := new(Stack)

	for i := len(s)-1; i >= 0; i-- {
		ch := s[i]
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {
			st.Push(string(ch))
		}else {
			top1 := st.Top()
			st.Pop()
			top2 := st.Top()
			st.Pop()

			new := "(" + top1 + string(ch) + top2 + ")"
			st.Push(new)
		}
	}
	return st.Top()
}

type Stack struct {
	items []string
}


func (s *Stack) Push(data string) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:s.Size()-1]
}

func (s *Stack) Top() string {
	if s.IsEmpty() {
		return ""
	}
	return s.items[s.Size()-1]
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}