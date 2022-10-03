package main

import "fmt"

func main() {
	// 类型定义
	type MyInt int
	var i MyInt
	i = 100
	fmt.Printf("type: %T val: %v\n", i, i)

	// 类型别名
	type MyInt2 = int
	var i2 MyInt2
	i2 = 200
	fmt.Printf("type: %T val: %v\n", i2, i2)

}
