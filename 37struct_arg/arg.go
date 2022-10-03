package main

import "fmt"

type Person struct {
	id   int
	name string
}

// 值传递，拷贝了一份副本
func changePerson(person Person) {
	person.id = 102
	person.name = "大宝"
	fmt.Printf("person: %v\n", person)
}

// 引用传递，借助指针进行操作
func changePerson2(person *Person) {
	person.id = 103
	person.name = "大大宝"
}

func main() {
	person := Person{
		101,
		"dabao",
	}
	fmt.Printf("person: %v\n", person)
	fmt.Println("--------")
	changePerson(person)
	fmt.Println("--------")
	fmt.Printf("person: %v\n", person)
	fmt.Println("--------")
	changePerson2(&person)
	fmt.Printf("person: %v\n", person)

}
