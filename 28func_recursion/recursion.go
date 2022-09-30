package main

import "fmt"

// 递归

func f1() {
	s := 1
	for i := 1; i <= 5; i++ {
		s *= i
	}
	fmt.Printf("s: %v\n", s)
}

func f2(a int) int {
	if a == 1 {
		return 1
	} else {
		return a * f2(a-1)
	}
}

// 斐波那契数列
// 计算公式：f(n) = f(n-1)+f(n-2) 且 f(1)=f(2)
func f(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return f(n-1) + f(n-2)
}
func main() {
	//f1()
	r := f2(5)
	fmt.Printf("r: %v\n", r)

	// 斐波那契数列
	r = f(5)
	fmt.Println(r)

}
