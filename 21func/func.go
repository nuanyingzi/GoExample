package main

import "fmt"

// 函数

// 求和
func sum(a, b int) (sum int) {
	return a + b
}

// 求两数最大值
func max(a, b int) (max int) {
	if a > b {
		return a
	} else {
		return b

	}
}

func test1() {
	a := 0.1 + 0.2 + 0.3 + 0.1 + 0.1 + 0.1 + 0.1
	b := 1.0
	if a == b {
		fmt.Println("==")
	} else {
		fmt.Println("<>")
	}
}

func main() {
	ret1 := sum(1, 2)
	fmt.Println(ret1)
	ret2 := max(1, 2)
	fmt.Println(ret2)
	test1()
}
