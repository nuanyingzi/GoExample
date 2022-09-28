package main

import "fmt"

// 函数类型

func sum(a, b int) int {
	return a + b
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	type f1 func(int, int) int

	var f1Sum f1
	f1Sum = sum
	r := f1Sum(1, 2)
	fmt.Printf("r: %v\n", r)

	var f1Max f1
	f1Max = max
	r = f1Max(91, 90)
	fmt.Printf("r: %v\n", r)

}
