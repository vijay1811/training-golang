package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
}

func modifyPerson(p person) {
	p.name = strings.Repeat(p.name, 10)
	fmt.Println(p)
}

func modifyPersonV1(p *person) {
	p.name = strings.Repeat(p.name, 10)
	fmt.Println(p)
}

func main() {
	p1 := person{"a", 10}
	fmt.Println(p1)
	p2 := person{
		name: "a",
		age:  10,
	}
	fmt.Println(p2)
	modifyPersonV1(&p2)
	fmt.Println(p2)

}
