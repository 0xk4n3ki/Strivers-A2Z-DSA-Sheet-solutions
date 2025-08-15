package main
import "fmt"

func main() {
	l, r := 3, 5
	res := rangeXor(l, r)
	fmt.Printf("l: %d, r: %d, res: %d\n", l, r, res)

	l, r = 1, 3
	res = rangeXor(l, r)
	fmt.Printf("l: %d, r: %d, res: %d\n", l, r, res)
}

func rangeXor(l, r int) int {
	lXor, rXor := xor(l-1), xor(r)
	return lXor ^ rXor
}

func xor(n int) int {
	switch n % 4 {
	case 0:
		return n
	case 1:
		return 1
	case 2:
		return n+1
	default:
		return 0
	}
}