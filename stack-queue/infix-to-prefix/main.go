package main

import (
	"fmt"
)

func main() {
	s := "(a+b)*c-d+f"
	fmt.Printf("s: %s\n", s)
	ans := toPrefix(s)
	fmt.Printf("ans: %s\n", ans)
}

func toPrefix(s string) string {
	ans := []rune{}
	st := new(Stack)

	s = revs(s)
	fmt.Printf("revs: %s\n", s)

	for _, ch := range s {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {
			ans = append(ans, ch)
		}else if ch == '(' {
			st.Push(ch)
		}else if ch == ')' {
			for !st.IsEmpty() && st.Top() != '(' {
				ans = append(ans, st.Top())
				st.Pop()
			}
			st.Pop()
		}else {
			if ch == '^' {
				for priority(ch) <= priority(st.Top()) {
					ans = append(ans, st.Top())
					st.Pop()
				}
			}else {
				for priority(ch) < priority(st.Top()) {
					ans = append(ans, st.Top())
					st.Pop()
				}
			}
			st.Push(ch)
		}
	}

	for !st.IsEmpty() {
		ans = append(ans, st.Top())
		st.Pop()
	}

	return revs(string(ans))
}

func priority(ch rune) int {
	switch ch {
	case '^' :
		return 3
	case '*', '/' :
		return 2
	case '+', '-' :
		return 1
	default:
		return -1
	}
}

func revs(s string) string {
	arr := []rune(s)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	for i := 0; i < len(s); i++ {
		if arr[i] == '(' {
			arr[i] = ')'
		} else if arr[i] == ')' {
			arr[i] = '('
		}
	}

	return string(arr)
}

type Stack struct {
	items []rune
}

func (s *Stack) Push(data rune) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:s.Size()-1]
}

func (s *Stack) Top() rune {
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
