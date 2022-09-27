package main

import "fmt"

// 没有返回值
func test1() {
	fmt.Println("我没有返回值")
}

// 一个返回值
func sum(a, b int) (ret int) {
	ret = a + b
	return ret
}

// 两个返回值
func test2() (name string, age int) {
	name = "tom"
	age = 19
	return name, age
}

func main() {
	//test1()
	sum := sum(1, 2)
	fmt.Println(sum)
	name, age := test2()
	fmt.Println(name, age)
}
