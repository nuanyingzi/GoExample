package main

import "fmt"

// 闭包基础
func add() func(y int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

// 闭包进阶
func cal(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	// 闭包基础
	f := add()
	r := f(10)
	fmt.Printf("r: %v\n", r)
	r = f(20)
	fmt.Printf("r: %v\n", r)

	f1 := add()
	r = f1(100)
	fmt.Printf("r: %v\n", r)
	r = f1(20)
	fmt.Printf("r: %v\n", r)

	// 闭包进阶
	c1, c2 := cal(100)
	fmt.Println(c1(1), c2(2))
	fmt.Println(c1(1), c2(2))
}
