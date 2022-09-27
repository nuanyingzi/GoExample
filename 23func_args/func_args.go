package main

import "fmt"

// 形参
func sum(a, b int) int {
	return a + b
}

// 值传递
func f1(a int) {
	a = 200
}

// 引用传递 会改变
func f2(s []int) {
	s[0] = 100
}

// 可变参数
func f3(args ...int) {
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}

func main() {

	r := sum(1, 2) // 实参
	fmt.Printf("r: %v\n", r)

	x := 100
	f1(x)
	fmt.Printf("x: %v\n", x)

	s1 := []int{1, 2, 3}
	f2(s1)
	fmt.Printf("s1: %v\n", s1)

	// 可变参数
	f3(1, 23, 3)
	f3(2, 3, 4, 5)
}
