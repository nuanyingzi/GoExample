package main

import "fmt"

type Pet interface {
	eat()
	sleep()
}

type Cat struct {
}

type Dog struct {
}

// Cat实现Pet接口

func (c Cat) eat() {
	fmt.Println("Cat eat...")
}

func (c Cat) sleep() {
	fmt.Println("Cat sleep...")
}

// Dog实现Pet接口

func (d Dog) eat() {
	fmt.Println("Dog eat...")
}

func (d Dog) sleep() {
	fmt.Println("Dog sleep...")
}

type Person struct {
}

// pet 既可以传递Dog也可以传递Cat
func (p Person) cure(pet Pet) {
	pet.eat()
	pet.sleep()
}

func main() {
	cat := Cat{}
	dog := Dog{}
	person := Person{}
	person.cure(cat)
	person.cure(dog)
}
