package main

import "fmt"

type Person struct {
	name string
}

// 值传入
func showPerson1(per Person) {
	per.name = "大大宝"
}

// 指针传入
func showPerson2(per *Person) {
	// 自动解引用
	per.name = "dadabao"
}

func (per Person) showPerson3() {
	per.name = "大大宝"
}

func (per *Person) showPerson4() {
	// 自动解引用
	per.name = "dadabao"
}

func main() {
	p1 := Person{
		name: "大宝",
	}
	fmt.Printf("p1: %v\n", p1)
	showPerson1(p1)
	fmt.Printf("p1: %v\n", p1)

	p2 := &Person{
		name: "dabao",
	}
	fmt.Printf("p2: %v\n", *p2)
	showPerson2(p2)
	fmt.Printf("p2: %v\n", *p2)

	fmt.Println("---------------")

	p1.showPerson3()
	fmt.Printf("p1: %v\n", p1)
	p2.showPerson4()
	fmt.Printf("p2: %v\n", *p2)

}
