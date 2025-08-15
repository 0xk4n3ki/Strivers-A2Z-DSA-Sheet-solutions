package main
import "fmt"

func main() {
	q := new(Queue)
	q.Push(372)
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)

	fmt.Println(q)

	fmt.Println(q.IsEmpty())

	fmt.Println(q.Top())

	q.Pop()

	fmt.Println(q)
}

type Queue struct {
	items []int
}

func (q *Queue) Push(data int) {
	q.items = append(q.items, data)
}

func (q *Queue) Pop() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("Queue is Empty")
	}

	tmp := q.items[0]
	q.items = q.items[1:]

	return tmp, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
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

