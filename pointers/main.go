package main

import "fmt"

func main() {
	i := 4
	p := &i
	*p = 20
	fmt.Println(i)
}
