package main

import "fmt"

func test1() {
	var a1 [3]int
	var a2 [6]string
	fmt.Printf("a1: %T, a1: %v\n", a1, a1)
	fmt.Printf("a2: %T, a2: %v\n", a2, a2)
}

func test2() {
	var a1 = [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(a1); i++ {
		fmt.Printf("a1[%v]: %v\n", i, a1[i])
	}
}
func test3() {
	var a = [...]int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Printf("a[%v]: %v\n", i, v)

	}
}

func main() {
	//test1()
	//test2()
	test3()
}
