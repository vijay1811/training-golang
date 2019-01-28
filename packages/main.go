package main

import (
	"fmt"
	"foo"
	"log"
)

func main() {
	fmt.Println("hello world")
	log.Println("hello world")
	fmt.Println(foo.PackageName)
}
