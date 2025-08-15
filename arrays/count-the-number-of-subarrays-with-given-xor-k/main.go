package main

import "fmt"

func main() {
	arr := []int {4, 2, 2, 6, 4}
	k := 6

	bruteforceApp(arr, k)
	betterApp(arr, k)
}



func betterApp(arr []int, target int) {
	xorMap := map[int]int {}
	count := 0
	xor := 0

	xorMap[0] = 1

	for i := 0; i < len(arr); i++ {
		xor ^= arr[i]

		// if xor == target {
		// 	count++
		// }else {
		// 	tmp := xor ^ target

		// 	count += xorMap[tmp]
		// }

		tmp := xor ^ target
		count += xorMap[tmp]

		xorMap[xor]++
	}

	fmt.Printf("count: %d\n", count)
}


func bruteforceApp(arr []int, target int) {
	count := 0

	for i := 0; i < len(arr); i++ {
		xor := 0
		for j := i; j < len(arr); j++ {
			xor ^= arr[j]
			if xor == target {
				count++
			}
		}
	}
	fmt.Printf("count: %d\n", count)
}