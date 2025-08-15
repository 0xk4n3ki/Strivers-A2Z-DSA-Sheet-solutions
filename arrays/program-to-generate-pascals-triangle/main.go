package main

import "fmt"

func main() {
	r, c := 5, 3
	nthElement(r, c)

	n := 6
	row := printRow(n)
	fmt.Printf("row: %v\n", row)

	printTriangle(6)
}

func printTriangle(n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("%v\n", printRow(i))
	}
}

func printRow(n int) []int {
	res := 1
	arr := []int {}
	arr = append(arr, 1)

	for i := 1; i < n; i++ {
		res = res*(n-i)/i

		arr = append(arr, res)
	}
	// fmt.Printf("row: %v\n", arr)
	
	return arr
}


// (n-1)C(c-1) find karna hai
func nthElement(r, c int) {
	res := 1
	for i := 0; i < c-1; i++ {
		res *= (r-i-1)  // because (r-1)C(c-1) find karna hai

		res /= (i+1)
	}
	fmt.Printf("resulting element: %d\n", res)
}