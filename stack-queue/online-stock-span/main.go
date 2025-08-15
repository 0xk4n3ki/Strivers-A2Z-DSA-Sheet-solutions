package main
import "fmt"

func main() {
	obj := Constructor()
	fmt.Println(obj.Next(28))
	fmt.Println(obj.Next(14))
	fmt.Println(obj.Next(28))
	fmt.Println(obj.Next(35))
	fmt.Println(obj.Next(46))
	fmt.Println(obj.Next(53))
	fmt.Println(obj.Next(66))
	fmt.Println(obj.Next(80))
	fmt.Println(obj.Next(87))
	fmt.Println(obj.Next(88))
}


type StockSpanner struct {
    items []int
	st Stack
}


func Constructor() StockSpanner {
    return StockSpanner{[]int {}, Stack{[]int {}}}
}


func (this *StockSpanner) Next(price int) int {
	i := len(this.items)
    this.items = append(this.items, price)

	for !this.st.IsEmpty() && this.items[this.st.Top()] <= price {
		this.st.Pop()
	}

	prev := -1
	if !this.st.IsEmpty() {
		prev = this.st.Top()
	}

	this.st.Push(i)
	return i - prev
}

type Stack struct {
	items []int
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() {
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) Top() int {
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}