package main
import "fmt"

func main() {
	x := float32(2)
	n := 10
	fmt.Printf("x: %f, n: %d, ans: %f\n", x, n, calc(x, n))

	x, n = 2.1, 3
	fmt.Printf("x: %f, n: %d, ans: %f\n", x, n, calc(x, n))

	x, n = 2, -2
	fmt.Printf("x: %f, n: %d, ans: %f\n", x, n, calc(x, n))
}

func calc(x float32, n int) float32 {
	ans := float32(1)

	sign := 1
	if n < 0 {
		sign = -1
		n = n * -1
	}

	for n > 0 {
		if n % 2 == 0 {
			x = x*x
			n /= 2
		}else {
			ans = ans * x
			n = n - 1
		}
	}

	if sign == -1 {
		return 1/float32(ans)
	}
	return float32(ans)
}