package main

import "fmt"

func test1() {
	var name string
	var pname *string
	name = "dabao"
	pname = &name
	fmt.Printf("name: %v\n", name)
	fmt.Printf("pname: %v\n", pname)
}

// 结构体指针
func test2() {
	type Person struct {
		id   int
		name string
		age  int
	}
	dabao := Person{
		101,
		"dabao",
		20,
	}
	var pDabao *Person

	pDabao = &dabao

	fmt.Printf("dabao: %v\n", dabao)
	fmt.Printf("p_dabao: %v\n", pDabao)
	fmt.Printf("p_dabao: %v\n", *pDabao)
}

func test3() {
	type Person struct {
		id   int
		name string
		age  int
	}
	var tom = new(Person)
	tom.id = 101
	tom.name = "大宝"
	tom.age = 21
	fmt.Printf("tom: %v\n", *tom)

}

func main() {
	//test2()
	test3()
}
