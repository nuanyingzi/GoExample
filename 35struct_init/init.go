package main

import "fmt"

func main() {
	type Person struct {
		id, age     int
		name, email string
	}
	var dabao Person
	dabao = Person{
		id:    101,
		age:   21,
		name:  "大宝",
		email: "dabao@gmail.com",
	}
	fmt.Printf("dabao: %v\n", dabao)
}
