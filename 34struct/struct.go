package main

// 结构体

import "fmt"

type Person struct {
	id    int
	name  string
	age   int
	email string
}

type Customer struct {
	id, age     int
	name, email string
}

func main() {
	var Tom Person
	fmt.Printf("Tom: %v\n", Tom)
	Tom.id = 1
	Tom.name = "Tom"
	Tom.age = 18
	Tom.email = "Tom@gmail.com"
	fmt.Printf("Tom: %v\n", Tom)

	var Kite Customer
	fmt.Printf("Kite: %v\n", Kite)

	// 匿名结构体
	var tom struct {
		id   int
		name string
		age  int
	}
	tom.id = 101
	tom.name = "Tom1"
	tom.age = 20
	fmt.Printf("tom: %v\n", tom)

}
