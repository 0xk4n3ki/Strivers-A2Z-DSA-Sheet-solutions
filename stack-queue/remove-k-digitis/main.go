package main

import (
	"fmt"
)

func main() {
	s := "33526221184202197273"
	k := 19
	ans := removeKdigits(s, k)
	fmt.Printf("ans: %s\n", ans)
}

func removeKdigits(s string, k int) string {
	st := new(Stack)
	for i := 0; i < len(s); i++ {
		for !st.IsEmpty() && k > 0 && (st.Top()-'0') > (s[i]-'0') {
			st.Pop()
			k--
		}
		st.Push(s[i])
		fmt.Printf("i: %d, s[i]: %d, st: %v\n", i, s[i], st)
	}

	for k > 0 {
		st.Pop()
		k--
	}
	if st.IsEmpty() {
		return "0"
	}

	ans := []byte {}
	for !st.IsEmpty() {
		ans = append(ans, st.Top())
		st.Pop()
	}
	for len(ans) != 0 && ans[len(ans)-1] == '0' {
		ans = ans[:len(ans)-1]
	}
	reverseS(ans)
	if len(ans) == 0 {
		return "0"
	}
	return string(ans)
}

func reverseS(arr []byte){
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}






type Stack struct {
	items []byte
}

func (s *Stack) Push(data byte) {
	s.items = append(s.items, data)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Top() byte {
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