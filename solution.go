package main

import "fmt"

// [OLD] https://takeuforward.org/data-structure/quick-sort-algorithm/
// [OLD] https://takeuforward.org/plus/dsa/problems/find-pairs-with-given-sum-in-doubly-linked-list
// [OLD] https://takeuforward.org/data-structure/segregate-even-and-odd-nodes-in-linkedlist
// [NEW] https://takeuforward.org/arrays/minimum-days-to-make-m-bouquets/
// [NEW] https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/description/
// [NEW] https://takeuforward.org/plus/dsa/problems/generate-binary-strings-without-consecutive-1s
// [NEW] https://takeuforward.org/data-structure/palindrome-partitioning/
// [NEW] https://leetcode.com/problems/powx-n/description/
// [NEW] https://takeuforward.org/data-structure/infix-to-prefix/
// [NEW] https://takeuforward.org/data-structure/sliding-window-maximum/
// [NEW] https://leetcode.com/problems/count-number-of-nice-subarrays/description/


func main() {
	arr := []int {1, 2, 4, 5, 6, 8, 9}
	target := 7
	head, back := arr2dll(arr)
	printdll(head)
	fmt.Printf("ans: %v\n", pair(head, back, target))
}

func pair(head, back *Node, target int) [][]int {
	next, prev := head, back
	ans := [][]int {}

	for next != prev {
		x, y := next.val, prev.val
		sum := x + y

		if sum <= target {
			if sum == target {
				ans = append(ans, []int {x, y})
			}
			next = next.front
		}else {
			prev = prev.back
		}
	}
	return ans
}

type Node struct {
	val int
	front *Node
	back *Node
}

func arr2dll(arr []int) (*Node, *Node) {
	head := &Node{0, nil, nil}
	curr := head

	for _, i := range arr {
		tmp := &Node{i, nil, curr}
		curr.front = tmp
		curr = tmp
	}

	head.front.back = nil
	return head.front, curr
}

func printdll(head *Node) {
	for head != nil {
		fmt.Printf("%p %v\n", head, head)
		head = head.front
	}
}