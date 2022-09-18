package main

import "fmt"

func main() {

	// 基本类型
	var name string = "Tom"
	age := 18
	b := false
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", b)

	// 派生类型
	a := 100
	p := &a
	fmt.Printf("%T\n", p)

	//
	s := []int{1, 23, 11}
	fmt.Printf("%T\n", s)

}
