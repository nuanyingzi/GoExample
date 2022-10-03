package main

import "fmt"

// 结构体的嵌套

type Dog struct {
	name  string
	color string
	age   int
}

type Person struct {
	dog  Dog
	name string
	age  int
}

func main() {
	dog := Dog{
		name:  "大黄",
		color: "yellow",
		age:   1,
	}
	per := Person{
		dog:  dog,
		name: "dabao",
		age:  21,
	}
	fmt.Printf("Person: %v\n", per)
}
