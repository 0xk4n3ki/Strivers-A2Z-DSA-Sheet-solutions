package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "a+b*(c^d-e)^(f+g*h)-i"
	ans := postfix(s)
	fmt.Printf("ans: %s\n", ans)
}

func postfix(s string) string {
	var ans strings.Builder
	st := new(Stack)

	for i:= 0; i < len(s); i++ {
		ch := s[i]
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {
			ans.WriteByte(ch)
		}else if ch == '(' {
			st.Push(ch)
		}else if ch == ')' {
			for !st.IsEmpty() && st.SafeTop() != '(' {
				ans.WriteByte(st.SafeTop())
				st.Pop()
			}
			st.Pop()
		}else {
			for !st.IsEmpty() && priority(ch) <= priority(st.SafeTop()) {
				ans.WriteByte(st.SafeTop())
				st.Pop()
			}
			st.Push(ch)
		}
	}

	for !st.IsEmpty() {
		ans.WriteByte(st.SafeTop())
		st.Pop()
	}
	return ans.String()
}

func priority(ch byte) int {
	switch ch {
	case '^' :
		return 3
	case '*', '/' :
		return 2
	case '+', '-' :
		return 1
	default :
		return -1
	}
}

type Stack struct {
	items []byte
}

func (s *Stack) Push(data byte) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() (byte, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	val := s.items[s.Size()-1]
	s.items = s.items[:s.Size()-1]
	return val, nil
}

func (s *Stack) Top() (byte, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	val := s.items[s.Size()-1]
	return val, nil
}

func (s *Stack) SafeTop() byte {
	val, err := s.Top()
	if err != nil {
		return 0
	}
	return val
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}