package main

import "fmt"

// 匿名函数

func main() {
	max := func(a, b int) int {
		return a + b
	}
	fmt.Println(max(1, 2))

	func(a, b int) {
		var sub int
		sub = a - b
		fmt.Printf("sub: %v\n", sub)
	}(2, 3)

}
