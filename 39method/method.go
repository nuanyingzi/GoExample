package main

import (
	"fmt"
)

type Person struct {
	name string
}

// (person)接收者 receiver
func (person Person) eat() {
	fmt.Printf("%v can eat\n", person.name)
}

func (person Person) sleep() {
	fmt.Printf("%v can sleep\n", person.name)
}

type Customer struct {
	name string
}

func (c Customer) login(pwd string) {
	fmt.Printf("%v 正在登录\n", c.name)
	fmt.Println("............")
	fmt.Println("............")
	fmt.Println("............")
	fmt.Println("............")
	if pwd == "123" {
		fmt.Println("您已登录成功")
	} else {
		fmt.Println("登录失败，请检查账号名或者密码")
	}
}

func main() {
	person := Person{name: "dabao"}
	person.eat()
	person.sleep()

	c := Customer{name: "dabao"}
	c.login("123")

}
