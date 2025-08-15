package main

import "fmt"

func main() {
	n := 10
	printRec(1, n)
}

func printRec(i, n int) {
	if i > n {
		return
	}
	fmt.Println(i)
	printRec(i+1, n)
}