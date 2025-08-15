package main
import "fmt"

func main() {
	arr := []int {10, 2, -5}
	ans := brute(arr)
	fmt.Printf("ans: %v\n", ans)

	arr = []int {2, -8, 5, 10, -7, 3, -4, 15}
	ans = brute(arr)
	fmt.Printf("ans: %v\n", ans)

	arr = []int {8, -8}
	ans = optimal(arr)
	fmt.Printf("ans: %v\n", ans)

	arr = []int {-2, -1, 1, 2}
	ans = optimal(arr)
	fmt.Printf("ans: %v\n", ans)
}


func optimal(arr []int) []int {
	st := new(Stack)

	for _, i := range arr {
		if i > 0 {
			st.Push(i)
		}else {
			for !st.IsEmpty() && st.Top() > 0 && st.Top() < i*-1 {
				st.Pop()
			}
			if !st.IsEmpty() && st.Top() == i*-1 {
				st.Pop()
			}else if st.IsEmpty() || st.Top() < 0 {
				st.Push(i)
			}
		}
	}
	ans := []int {}

	for !st.IsEmpty() {
		ans = append(ans, st.Top())
		st.Pop()
	}

	reverseArr(ans)
	return ans
}



func brute(arr []int) []int {
	st := new(Stack)
	ans := []int {}

	for j := len(arr)-1; j >= 0; j-- {
		i := arr[j]
		if i < 0 {
			st.Push(i)
		}else {
			for !st.IsEmpty() {
				if st.Top()*-1 > i {
					break
				}else if st.Top()*-1 < i {
					st.Pop()
				}else {
					break
				}
			}
			if st.IsEmpty() {
				ans = append(ans, i)
			}else {
				if st.Top()*-1 == i {
					st.Pop()
				}
			}
		}
	}
	neg := []int {}
	for !st.IsEmpty() {
		neg = append(neg, st.Top())
		st.Pop()
	}

	reverseArr(ans)

	neg = append(neg, ans...)

	return neg
}

func reverseArr(arr []int) {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
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