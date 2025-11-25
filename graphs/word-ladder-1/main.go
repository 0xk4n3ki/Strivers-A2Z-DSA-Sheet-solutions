package main

import (
	"fmt"
)

func main() {
	wordlist := []string {"des", "der", "dfr", "dgt", "dfs"}
	startword, targetword := "der", "dfs"
	fmt.Printf("ans: %d\n", calc(wordlist, startword, targetword))

	wordlist = []string {"geek", "gefk"}
	startword, targetword = "gedk", "geek"
	fmt.Printf("ans: %d\n", calc(wordlist, startword, targetword))
}

func calc(wordlist []string, startword, targetword string) int {
	q := new(Queue)
	q.Push(startword, 1)

	set := map[string]struct{} {}
	for _, i := range wordlist {
		set[i] = struct{}{}
	}
	delete(set, startword)

	for !q.IsEmpty() {
		word, steps := q.Top().word, q.Top().steps
		q.Pop()

		if word == targetword {
			return steps
		}

		wArr := []rune(word)
		for i := 0; i < len(wArr); i++ {
			original := wArr[i]
			for ch := 'a'; ch <= 'z'; ch++ {
				wArr[i] = ch
				newWord := string(wArr)

				if _, ok := set[newWord]; ok {
					delete(set, newWord)
					q.Push(newWord, steps+1)
				}
			}
			wArr[i] = original
		}
	}
	return 0
}


type Pair struct {
	word string
	steps int
}

type Queue struct {
	Arr []Pair
}

func (q *Queue) Push(word string, steps int) {
	q.Arr = append(q.Arr, Pair{word, steps})
}

func (q *Queue) Size() int {
	return len(q.Arr)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Top() Pair {
	if q.IsEmpty() {
		return Pair{}
	}
	return q.Arr[0]
}

func (q *Queue) Pop() {
	if q.IsEmpty() {
		return
	}
	q.Arr = q.Arr[1:]
}