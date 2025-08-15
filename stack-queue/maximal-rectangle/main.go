package main
import "fmt"
func main() {
	matrix := [][]byte {
		{'1','0','1','0','0'},
		{'1','0','1','1','1'},
		{'1','1','1','1','1'},
		{'1','0','0','1','0'},
	}
	ans := maximal(matrix)
	fmt.Printf("ans: %d\n", ans)
}

func maximalRectangle(matrix [][]byte) int {
	m,n := len(matrix), len(matrix[0])
	newMat := make([][]int, m)
	for i := 0; i < m; i++ {
		newMat[i] = make([]int, n)
	}

	for j := 0; j < n; j++ {
		sum := 0
		for i := 0; i < m; i++ {
			if matrix[i][j] == '1' {
				sum++
			}else {
				sum = 0
			}
			newMat[i][j] = sum
		}
	}
	max := 0
	for i := 0; i < m; i++ {
		area := largestRectangle(newMat[i])
		if area > max {
			max = area
		}
	}
	return max
}

func largestRectangle(arr []int) int {
	st := new(Stack)
	n := len(arr)
	max := 0

	for i := 0; i < n; i++ {
		for !st.IsEmpty() && arr[st.Top()] > arr[i] {
			height := arr[st.Top()]
			st.Pop()
			width := (i-st.Top()-1)
			area := height * width
			if area > max {
				max = area
			}
		}
		st.Push(i)
	}

	for !st.IsEmpty() {
		height := arr[st.Top()]
		st.Pop()
		width := (n-st.Top()-1)
		area := height * width
		if area > max {
			max = area
		}
	}
	return max
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