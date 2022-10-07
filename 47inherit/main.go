package main

import "fmt"

// 继承

type Animal struct {
	name string
	age  int
}

func (a Animal) eat() {
	fmt.Println("Eat ...")

}

func (a Animal) sleep() {
	fmt.Println("Sleep ...")
}

type Dog struct {
	Animal
	color string
}

type Cat struct {
	Animal
	skill string
}

func main() {
	dog := Dog{
		Animal: Animal{
			name: "犬类",
			age:  8,
		},
		color: "yellow",
	}
	dog.eat()
	dog.sleep()

	cat := Cat{
		Animal{
			"猫科",
			2,
		},
		"climb",
	}
	cat.eat()
	cat.sleep()
}
