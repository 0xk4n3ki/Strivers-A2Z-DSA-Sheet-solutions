package main
import "fmt"

type Node struct {
	val int
	next *Node
	back *Node
}

func main() {
	arr := []int {1, 2, 4, 5, 6, 8, 9}
	head := arr2dll(arr)
	printdll(head)

	findPairs(head, 7)
	usingTwoPointers(head, 7)

	arr = []int {1, 5, 6}
	head = arr2dll(arr)
	printdll(head)

	findPairs(head, 6)
	usingTwoPointers(head, 6)
}


func usingTwoPointers(head *Node, target int) {
	if head == nil || head.next == nil {
		fmt.Printf("ans: -1\n")
	}

	left := head
	right := head
	for right.next != nil {
		right = right.next
	}

	ans := [][]int {}

	for left != right {
		sum := left.val + right.val

		if sum == target {
			ans = append(ans, []int {left.val, right.val})
			left = left.next
			right = right.back
		}else if sum < target {
			left = left.next
		}else {
			right = right.back
		}

		if left == right || right.next == left {
			break
		}
	}
	fmt.Printf("[twoPointers] ans: %v\n", ans)
}



func findPairs(head *Node, target int) {
	curr := head
	eleMap := map[int]struct{} {}
	ans := [][]int {}

	for curr != nil {
		rem := target - curr.val

		if _, ok := eleMap[rem]; ok {
			if curr.val < rem {
				ans = append(ans, []int {curr.val, rem})
			}else {
				ans = append(ans, []int {rem, curr.val})
			}
		}
		eleMap[curr.val] = struct{}{}

		curr = curr.next
	}
	fmt.Printf("[using map] ans: %v\n", ans)
}









func arr2dll(arr []int) *Node {
	head := &Node{arr[0], nil, nil}

	prev := head

	for i := 1; i < len(arr); i++ {
		temp := &Node{arr[i], nil, prev}

		prev.next = temp
		prev = temp
	}

	return head
}


func printdll(head *Node) {
	for head != nil {
		fmt.Printf("%d", head.val)

		if head.next != nil {
			fmt.Printf(" <-> ")
		}
		head = head.next
	}
	fmt.Println()
}