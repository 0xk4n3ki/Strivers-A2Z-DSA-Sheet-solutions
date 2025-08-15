package main

import "fmt"

func main() {
	n := 10

	printRec(n)
}

func printRec(n int) {
	if n == 0 {
		return 
	}
	fmt.Println(n)
	printRec(n-1)
}