package main

import "fmt"

func language(lang string) {
	switch lang {
	case "java":
		fmt.Println("lang: java")
	case "go":
		fmt.Println("lang: go")
	case "python":
		fmt.Println("lang: python")
	default:
		fmt.Println("lang: " + lang + "\r")
	}
}

func add(a, b int) int {
	return a + b
}

func rangeOfNum(n int) {
	switch {
	case n < 10:
		fmt.Println("range<10")
	case n < 20:
		fmt.Println("10<= range<20")
	default:
		fmt.Println("range>=20")
	}
}

func main() {
	language("ruby")
}
