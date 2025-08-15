package main

import "fmt"

func main() {
	dup := 123456
	input := dup
	rev := 0
	
	for input > 0 {
		ones := input % 10
		rev = rev * 10 + ones
		input = input / 10
	}

	fmt.Printf("reverse of %d : %d\n", dup, rev)
}