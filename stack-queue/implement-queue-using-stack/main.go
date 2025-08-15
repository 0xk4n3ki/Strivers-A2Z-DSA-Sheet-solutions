package main
import "fmt"

func main() {
	q := new(Queue)
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)

	fmt.Println(q)

	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	fmt.Println(q)
}

type Queue struct {
	input, output Stack
}

func (q *Queue) Push(data int) {
	q.input.Push(data)
}

func (q *Queue) Pop() (int, error) {
	if !q.output.IsEmpty() {
		return q.output.Pop()
	}else {
		for !q.input.IsEmpty() {
			val, _ := q.input.Pop()
			q.output.Push(val)
		}
		return q.output.Pop()
	}
}

func (q *Queue) Top() (int, error) {
	if !q.output.IsEmpty() {
		return q.output.Top()
	}else {
		for !q.input.IsEmpty() {
			val, _ := q.input.Pop()
			q.output.Push(val)
		}
		return q.output.Top()
	}
}

func (q *Queue) IsEmpty() bool {
	return (q.output.IsEmpty() && q.input.IsEmpty())
}



// type Queue struct {
// 	stk Stack
// }

// func (q *Queue) Push(data int) {
// 	stk2 := Stack{[]int {}}
// 	for !q.stk.IsEmpty() {
// 		val, _ := q.stk.Pop()
// 		stk2.Push(val)
// 	}
	
// 	stk2.Push(data)

// 	for !stk2.IsEmpty() {
// 		val, _ := stk2.Pop()
// 		q.stk.Push(val)
// 	}
// }

// func (q *Queue) Pop() (int, error) {
// 	return q.stk.Pop()
// }

// func (q *Queue) IsEmpty() bool {
// 	return q.stk.IsEmpty()
// }

// func (q *Queue) Size() int {
// 	return q.stk.Size()
// }





type Stack struct {
	items []int
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() (int, error) {
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

func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) Size() int {
	return len(s.items)
}
