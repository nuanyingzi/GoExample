package main

import "fmt"

// 高阶函数

func sayHello(name string) {
	fmt.Printf("Hello %v\n", name)
}

func test(name string, f func(string)) {
	f(name)
}

// 函数作为返回值

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func cal(operate string) func(int, int) int {
	switch operate {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {
	test("Tom", sayHello)

	ff := cal("+")
	r := ff(1, 2)
	fmt.Println(r)

	ff2 := cal("-")
	r = ff2(1, 2)
	fmt.Println(r)
}
