package main

import (
	"fmt"
	"math"
)

func sumOfNumbers(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func sumOfNumbersV1(n int) int {
	sum := 0
	i := 1
	for i <= n {
		sum += i
		i++
	}
	return sum
}

func nestedForWithBreak(n int) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			fmt.Println(i, j)
		}
	}
}

func forever() {
	i := 0

	for i < 100 {
		for {
			if i == math.MaxInt8 {
				fmt.Println("forever exits")
				break
			}
			i++
		}
	}
}

func main() {
	fmt.Println(sumOfNumbersV1(2))
	forever()
}
