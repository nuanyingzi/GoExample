package main

import "fmt"

type Pet interface {
	eat(string) string
}

type Dog struct {
	name string
}

//func (d Dog) eat(name string) string {
//	d.name = "花花..."
//	fmt.Printf("食物: %v\n", name)
//	return "吃的还行"
//}

func (d *Dog) eat2(name string) string {
	d.name = "花花..."
	fmt.Printf("食物: %v\n", name)
	return "吃的还行"
}

func main() {
	/*dog := Dog{name: "花花"}
	s := dog.eat("shit")
	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", dog)*/

	dog := &Dog{name: "花花"}
	s := dog.eat2("shit")
	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", *dog)

}
