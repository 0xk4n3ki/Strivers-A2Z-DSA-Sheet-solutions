package main


func main() {

}

type LRUCache struct {
    head, tail *Node
	hashMap map[int]*Node
	capacity int
}

type Node struct {
	key, value int
	next *Node
	back *Node
}


func Constructor(capacity int) LRUCache {
	head := &Node{-1, -1, nil, nil}
	tail := &Node{-1, -1, nil, head}
	head.next = tail
    return LRUCache{
		head: head,
		tail: tail,
		hashMap: map[int]*Node {},
		capacity: capacity,
	}
}


func (this *LRUCache) Get(key int) int {
    if _, ok := this.hashMap[key]; !ok {
		return -1
	}
	curr := this.hashMap[key]
	deleteNode(curr)
	insertAfterHead(this.head, curr)

	return curr.value
}


func (this *LRUCache) Put(key int, value int)  {
    if _, ok := this.hashMap[key]; ok {
		curr := this.hashMap[key]
		curr.value = value
		deleteNode(curr)
		insertAfterHead(this.head, curr)
	}else {
		if len(this.hashMap) == this.capacity {
			toDelete := this.tail.back
			delete(this.hashMap, toDelete.key)
			deleteNode(toDelete)
		}
		newNode := &Node{key: key, value: value, next: nil, back: nil}
		this.hashMap[key] = newNode
		insertAfterHead(this.head, newNode)
	}
}

func deleteNode(curr *Node) {
	prevNode := curr.back
	nextNode := curr.next

	prevNode.next = nextNode
	nextNode.back = prevNode
}

func insertAfterHead(head, newNode *Node) {
	currAfterHead := head.next
	head.next = newNode
	newNode.back = head
	newNode.next = currAfterHead
	currAfterHead.back = newNode
}