package main

import "fmt"

func main() {
	arr := []int {5, 5, 5, 10, 20}
	fmt.Println(check(arr))

	arr = []int {5, 5, 10, 10, 20}
	fmt.Println(check(arr))
}

func check(arr []int) bool {
	fives, tens := 0, 0

	for _, i := range arr {
		rem := i - 5

		if rem == 5 {
			tens++
			if fives > 0 {
				fives--
			}else {
				fmt.Println("not enough fives for $10 change")
				return false
			}
		}else if rem == 15 {
			if tens > 0 && fives > 0 {
				tens--
				fives--
			}else if fives > 2 {
				fives -= 3
			}else {
				fmt.Println("not enough change for $20")
				return false
			}
		}else {
			fives++
		}
	}
	return true
}