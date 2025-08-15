package main

import "fmt"

func main() {
	n := 5
	name(n)
}

func name(n int) {
	if n == 0{
		return
	}
	fmt.Println("suraj yadav")
	n--
	name(n)
}