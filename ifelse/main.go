package main

import "fmt"

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}

	return false
}

func isEvenV1(n int) bool {
	if q := isEven(n); q {
		return true
	}

	return false
}

func isEvenV2(n int) bool {
	if q := n % 2; q%2 == 0 {
		return true
	} else {
		return false
	}
}

func rangeOfNum(n int) {
	if n < 10 {
		fmt.Println("range < 10")
	} else if n < 20 {
		fmt.Println("10<= range < 20")
	} else {
		fmt.Println("range > 20")
	}
}

// func ifElse(n int) {
// 	if n%2 == 0 {
// 		fmt.Println("number is divisible by two")
// 	}else
// }

func main() {

}
