package main

import "fmt"

//var i, j, k = 3 + 2i, "hi", 2.2

func main() {
	i, j, k := 3+2i, "hi", 2.2
	fmt.Println(i, j, k)
	fmt.Printf("%T,%T,%T\n", i, j, k)
}
