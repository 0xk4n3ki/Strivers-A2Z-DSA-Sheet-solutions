package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "tree"
	try(s)
	frequencySort(s)
	fSort(s)

	s = "Aabb"
	try(s)
	frequencySort(s)
	fSort(s)
}

func fSort(s string) {
	arr := make([]string, 256)
	for _, c := range s {
		arr[c] += string(c)
	}

	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) > len(arr[j])
	})

	var builder strings.Builder

	for _, m := range arr {
		builder.WriteString(m)
	}
	fmt.Printf("s: %s, ans: %s\n", s, builder.String())
}

func try(s string) {
	freq := map[int32]int {}

	for _, char := range s {
		// fmt.Printf("%c, %T\n", char, char)
		freq[char]++
	}
	// fmt.Printf("freq: %v\n", freq)

	type kv struct {
		key int32
		value int
	}

	var sorted []kv
	for k, v := range freq {
		sorted = append(sorted, kv{k, v})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].value > sorted[j].value
	})

	// fmt.Printf("sorted: %v\n", sorted)

	var builder strings.Builder

	for _, j := range sorted {
		for i := j.value; i > 0; i-- {
			builder.WriteRune(j.key)
		}
	}
	fmt.Printf("s: %s, ans: %s\n", s, builder.String())
}


func frequencySort(s string) string {
	freq := make([]int, 255)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		if freq[b[i]] == freq[b[j]] {
			return b[i] > b[j]
		}
		return freq[b[i]] > freq[b[j]]
	})

	fmt.Printf("s: %s, ans: %s\n", s, string(b))
	return string(b)
}