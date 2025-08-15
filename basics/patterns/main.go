package main

import "fmt"

func main() {
	fmt.Println("======Rectangular Star Pattern======")
	rectangularStarPattern(6)

	fmt.Println("======Right Angled Triangle======")
	rightAngledTriangle(6)

	fmt.Println("======Right Angled Number Pyramid======")
	rightAngledNumberPyramid(6)

	fmt.Println("======Right Angled Number Pyramid 2======")
	rightAngledNumberPyramid2(6)

	fmt.Println("======Inverted Right Pyramid======")
	invertedRightPyramid(6)

	fmt.Println("======Inverted Numbered Right Pyramid======")
	invertedNumberedRightPyramid(6)

	fmt.Println("======Star Pyramid======")
	starPyramid(6)

	fmt.Println("======Inverted Star Pyramid======")
	invertedStarPyramid(6)

	fmt.Println("======Diamond Star Pattern======")
	diamondStarPattern(6)

	fmt.Println("======Half Diamond Star Pattern======")
	halfDiamondStarPattern(6)

	fmt.Println("======Binary Number Triangle Pattern======")
	binaryNumberTrianglePattern(6)

	fmt.Println("======Number Crown Pattern======")
	numberCrownPattern(6)

	fmt.Println("======Increasing Number Triangle Pattern======")
	increasingNumberTrianglePattern(6)

	fmt.Println("======Increasing Letter Triangle Pattern======")
	increasingLetterTrianglePattern(6)

	fmt.Println("======Reverse Letter Triangle Pattern======")
	reverseLetterTrianglePattern(6)

	fmt.Println("======Alpha-Ramp Pattern======")
	alphaRampPattern(6)

	fmt.Println("======Alpha-Hill Pattern======")
	alphaHillPattern(6)

	fmt.Println("======Alpha-Triangle Pattern======")
	alphaTrianglePattern(6)

	fmt.Println("======Symmetric-Void Pattern======")
	symmetricVoidPattern(6)

	fmt.Println("======Symmetric-Butterfly Pattern======")
	symmetricButterflyPattern(6)

	fmt.Println("======Hollow Rectangle Pattern======")
	hollowRectanglePattern(6)

	fmt.Println("======The Number Pattern======")
	theNumberPattern(6)
}

func rectangularStarPattern(a int) {
	for i := 0; i < a; i++ {
		for j := 0; j < a; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func rightAngledTriangle(a int) {
	for i := 0; i < a; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func rightAngledNumberPyramid(a int) {
	for i := 0; i < a; i++ {
		for j := 1; j <= i+1; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}

func rightAngledNumberPyramid2(a int) {
	for i := 0; i < a; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d ", i+1)
		}
		fmt.Println()
	}
}

func invertedRightPyramid(a int) {
	for i := a; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func invertedNumberedRightPyramid(a int) {
	for i := a; i > 0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}

func starPyramid(a int) {
	for i := 1; i <= a; i++ {
		for j := 0; j < a-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func invertedStarPyramid(a int) {
	for i := 0; i < a; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*(a-i)-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func diamondStarPattern(a int) {
	for i := 0; i < a; i++ {
		for j := 0; j < a-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := a; i > 0; i-- {
		for j := 0; j < a-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func halfDiamondStarPattern(a int) {
	// for i := 0; i < a; i++ {
	// 	for j := 0; j < i+1; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }
	// for i := a-1; i > 0; i-- {
	// 	for j := 0; j < i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	for i := 0; i < 2*a-1; i++ {
		stars := i+1
		if i >= a {
			stars = 2*a - (i+1)
		}
		for j := 0; j < stars; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func binaryNumberTrianglePattern(a int) {
	b := 0
	for i := 0; i < a; i++ {
		if i % 2 == 0 {
			b = b ^ 1
		}
		for j := 0; j < i+1; j++ {
			fmt.Printf("%d", b)
			b = b ^ 1
		}
		fmt.Println()
	}
}

func numberCrownPattern(a int) {
	for i := 0; i < a; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("%d", j+1)
		}
		for j := 0; j < 2*(a-i-1); j++ {
			fmt.Print(" ")
		}
		for j := i+1; j > 0; j-- {
			fmt.Printf("%d", j)
		}
		fmt.Println()
	}
}

func increasingNumberTrianglePattern(a int) {
	count := 1
	for i := 0; i < a; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("%d ", count)
			count++
		}
		fmt.Println()
	}
}

func increasingLetterTrianglePattern(a int) {
	for i := 0; i < a; i++{
		char := 'A'
		for j := 0; j < i+1; j++ {
			fmt.Printf("%c ", char)
			char++
		}
		fmt.Println()
	}
	
}

func reverseLetterTrianglePattern(a int) {
	for i := a; i > 0; i-- {
		char := 'A'
		for j := 0; j < i; j++ {
			fmt.Printf("%c ", char)
			char++
		}
		fmt.Println()
	}
}

func alphaRampPattern(a int) {
	char := 'A'
	for i := 0; i < a; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("%c ", char)
		}
		fmt.Println()
		char++
	}
}

func alphaHillPattern(a int) {
	for i := 0; i < a; i++ {
		char := 'A'
		for j := 0; j < a-i-1; j++{
			fmt.Print(" ")
		}
		for j := 0; j < i+1; j++ {
			fmt.Printf("%c", char)
			char++
		}
		char--
		for j := 0; j < i; j++ {
			char--
			fmt.Printf("%c", char)
		}
		fmt.Println()
	}
}

func alphaTrianglePattern(a int) {
	char := 'A' + a - 1
	for i := 0; i < a; i++ {
		c := char - i
		for j := 0; j < i+1; j++ {
			fmt.Printf("%c", c)
			c++
		}
		fmt.Println()
	}
}

func symmetricVoidPattern(a int) {
	for i := 0; i < a; i++ {
		for j := a-i; j > 0; j-- {
			fmt.Print("*")
		}
		for j := 0; j < 2*i; j++ {
			fmt.Print(" ")
		}
		for j := a-i; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 0; i < a; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		for j := 2*(a-i-1); j>=0; j-- {
			fmt.Print(" ")
		}
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func symmetricButterflyPattern(a int) {
	// for i := 0; i < a; i++ {
	// 	for j := 0; j < i+1; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	for j := 0; j < 2*(a-i-1); j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for j := 0; j < i+1; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }
	// for i := a-1; i > 0; i-- {
	// 	for j := i; j > 0; j-- {
	// 		fmt.Print("*")
	// 	}
	// 	for j := 2*(a-i); j > 0; j-- {
	// 		fmt.Print(" ")
	// 	}
	// 	for j := i; j > 0; j-- {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }


	for i := 1; i <= 2*a-1; i++ {
		var stars, spaces int
		if i <= a{
			stars = i
			spaces = 2*(a-i)
		}else{
			stars = 2*a - i
			spaces = 2*(i-a)
		}
		for j := 0; j < stars; j++ {
			fmt.Print("*")
		}
		for j := 0; j < spaces; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < stars; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func hollowRectanglePattern(a int) {
	// for i := 0; i < a; i++ {
	// 	if i == 0 || i == a-1 {
	// 		for j := 0; j < a; j++ {
	// 			fmt.Print("*")
	// 		}
	// 	}else{
	// 		fmt.Print("*")
	// 		for j := 0; j < a-2; j++ {
	// 			fmt.Print(" ")
	// 		}
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	for i := 0; i < a; i++ {
		for j := 0; j < a; j++ {
			if i == 0 || i == a-1 || j == 0 || j == a-1 {
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func theNumberPattern(a int) {
	for i := 0; i < 2*a-1; i++ {
		for j := 0; j < 2*a-1; j++ {
			top := i
			left := j
			bottom := 2*(a-1)-i
			right := 2*(a-1)-j
			fmt.Printf("%d ", a-min(min(top, bottom), min(left, right)))
		}
		fmt.Println()
	}
}