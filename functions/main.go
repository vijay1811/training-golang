package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// func addAll(params ...int) int {
// 	if len(params)==1{
// 		return params[0]
// 	}
// 	return x + y
// }

func swap(x, y int) (int, int) {
	return y, x
}

func divide(divident, divisor int) (quotient int, remainder int) {
	quotient = divident / divisor
	if divident%2 != 0 {
		quotient := 0
		remainder = divident % divisor
		return quotient, remainder
	}
	return
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(swap(2, 1))
	fmt.Println(divide(5, 4))
}
