package main

// 变量

import "fmt"

func getNameAndAge() (string, int) {
	return "大宝", 18
}

func main() {

	// 初始化 赋初值
	var name string = "dabao"
	var age uint8
	var married bool

	fmt.Printf("name = %v, type = %T\n", name, name)
	fmt.Printf("age = %v, type = %T\n", age, age)
	fmt.Printf("married = %v, type = %T\n", married, married)

	// 类型推断
	var s = "05string"
	fmt.Printf("s = %v, type = %T\n", s, s)

	// 短变量声明
	shi := "是"
	fmt.Printf("shi = %v, type = %T\n", shi, shi)

	// 匿名变量
	dName, dAge := getNameAndAge()
	fmt.Println(dName, dAge)

	dName1, _ := getNameAndAge()
	fmt.Println(dName1)

}
