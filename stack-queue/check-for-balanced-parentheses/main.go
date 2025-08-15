package main

import "fmt"

func main() {
	s := "()[{}()]"
	ans := check(s)
	fmt.Printf("ans: %t\n", ans)

	fmt.Println(check("()[{}()]")) // true
	fmt.Println(check("()[{()}]")) // true
	fmt.Println(check("()[{()]"))  // false
	fmt.Println(check("["))        // false
	fmt.Println(check("([]{})"))   // true
	fmt.Println(check("([]{)}"))   // false

}

func check(s string) bool {
	st := new(Stack)
	for _, ch := range s {
		switch ch {
		case '(', '{', '[':
			st.Push(ch)
		case ')', '}', ']':
			char, err := st.Pop()
			if err != nil {
				fmt.Printf("error: %v\n", err)
				return false
			}
			if (ch == ')' && char != '(') ||
				(ch == '}' && char != '{') ||
				(ch == ']' && char != '[') {
				fmt.Printf("char: %c, ch: %c\n", char, ch)
				return false
			}
		}
	}
	return st.Size() == 0
}

type Stack struct {
	items []rune
}

func (s *Stack) Push(data rune) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() (rune, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Top() (rune, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	val := s.items[len(s.items)-1]
	return val, nil
}

func (s *Stack) Size() int {
	return len(s.items)
}
