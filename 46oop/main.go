package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) eat() {
	fmt.Println("Eat...")
}

func (p Person) sleep() {
	fmt.Println("sleep...")
}

func (p Person) work() {
	fmt.Println("work...")
}

func main() {
	person := Person{
		Name: "dabao",
		Age:  22,
	}
	person.eat()
	person.sleep()
	person.work()
}
