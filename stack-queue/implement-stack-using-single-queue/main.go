package main
import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) push(data int) {
	q.items = append(q.items, data)
}

func (q *Queue) Pop() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("Queue is empty")
	}
	top := q.items[0]
	q.items = q.items[1:]
	return top, nil
}

func (q *Queue) Top() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("Queue is empty")
	}
	return q.items[0], nil
}

func (q *Queue) Size() int {
	return len(q.items)
}

func main() {
	s := new(Stack)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println(s)

	fmt.Println(s.Pop())
	fmt.Println(s)
	fmt.Println(s.Top())
}

type Stack struct {
	qe Queue
}

func (s *Stack) Push(data int) {
	s.qe.push(data)
	n := s.qe.Size()

	for i := 0; i < n-1; i++ {
		val, _ := s.qe.Pop()
		s.qe.push(val)
	}
}

func (s *Stack) Pop() (int, error) {
	return s.qe.Pop()
}

func (s *Stack) Top() (int, error) {
	return s.qe.Top()
}

func (s *Stack) Size() int {
	return s.qe.Size()
}
