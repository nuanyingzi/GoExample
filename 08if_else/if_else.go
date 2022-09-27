package main

import "fmt"

func f2() {
	var num int
	fmt.Println("请输入一个整数：")
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Println("输入的的是一个偶数")
	} else {
		fmt.Println("奇数")
	}
}

func f3() {
	if score := 88; score >= 80 {
		fmt.Println("优秀")
	} else if score < 80 && score >= 60 {
		fmt.Println("良好")
	} else {
		fmt.Println("不及格")
	}
}

func main() {
	/*var (
		name  05string
		age   int
		email 05string
	)
	06fmt.Println("请输入name,age,email, 用空格分割")
	06fmt.Scan(&name, &age, &email)
	06fmt.Printf("name: %v\n", name)
	06fmt.Printf("age: %v\n", age)
	06fmt.Printf("email: %v\n", email)*/

	//f2()

	f3()
}
