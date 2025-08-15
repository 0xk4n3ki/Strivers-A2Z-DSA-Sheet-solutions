package main
import "fmt"

func main() {
	a, b := 5, 10
	fmt.Printf("a: %d, b: %d, ", a, b)
	
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("a: %d, b: %d, ", a, b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("a: %d, b: %d\n", a, b)
}