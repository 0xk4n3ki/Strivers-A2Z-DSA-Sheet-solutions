package main

func main() {

}

type MinStack struct {
	items []int
	min   int
}

func Constructor() MinStack {
	return MinStack{[]int{}, -1}
}

func (this *MinStack) Push(val int) {
	if this.IsEmpty() {
		this.items = append(this.items, val)
		this.min = val
	} else {
		if this.min > val {
			this.items = append(this.items, 2*val-this.min)
			this.min = val
		} else {
			this.items = append(this.items, val)
		}
	}
}

func (this *MinStack) IsEmpty() bool {
	return len(this.items) == 0
}

func (this *MinStack) Size() int {
	return len(this.items)
}

func (this *MinStack) Pop() {
	if this.IsEmpty() {
		return
	}
	if this.items[this.Size()-1] < this.min {
		this.min = 2*this.min - this.items[this.Size()-1]
	}
	this.items = this.items[:this.Size()-1]
}

func (this *MinStack) Top() int {
	if this.IsEmpty() {
		return -1
	}
	if this.items[this.Size()-1] < this.min {
		return this.min
	}
	return this.items[this.Size()-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

// --------------------- O(n) time complexity in Pop ------------------------

// type MinStack struct {
//     items []int
// 	min int
// }

// func Constructor() MinStack {
//     return MinStack{[]int {}, -1}
// }

// func (this *MinStack) Push(val int)  {
//     this.items = append(this.items, val)
// 	if this.min == -1 {
// 		this.min = this.Size()-1
// 	}else {
// 		if this.items[this.min] > val {
// 			this.min = this.Size()-1
// 		}
// 	}
// }

// func (this *MinStack) Pop()  {
//     if this.IsEmpty() {
// 		return
// 	}
// 	if this.min == this.Size()-1 {
// 		x := 0
// 		for i := 0; i < this.Size()-1; i++ {
// 			if this.items[x] > this.items[i] {
// 				x = i
// 			}
// 		}
// 		this.min = x
// 	}
// 	this.items = this.items[:this.Size()-1]
// }

// func (this *MinStack) Top() int {
//     if this.IsEmpty() {
// 		return -1
// 	}
// 	return this.items[this.Size()-1]
// }

// func (this *MinStack) GetMin() int {
// 	return this.items[this.min]
// }

// func (this *MinStack) Size() int {
// 	return len(this.items)
// }

// func (this *MinStack) IsEmpty() bool {
// 	return this.Size() == 0
// }
