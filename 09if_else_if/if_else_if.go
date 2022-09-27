package main

import "fmt"

func f1() {
	// Monday Tuesday Wednesday Thursday Friday Saturday Sunday
	var s string
	fmt.Println("请输入第一个字符：")
	fmt.Scan(&s)

	if s == "M" {
		fmt.Println("星期一")
	} else if s == "T" {
		fmt.Println("请输入第二个字符：")
		fmt.Scan(&s)
		if s == "u" {
			fmt.Println("星期二")
		} else if s == "h" {
			fmt.Println("星期四")
		} else {
			fmt.Println("输入错误")
		}
	} else if s == "W" {
		fmt.Println("星期三")
	} else if s == "F" {
		fmt.Println("星期五")
	} else if s == "S" {
		fmt.Println("请输入第二个字符：")
		fmt.Scan(&s)
		if s == "a" {
			fmt.Println("星期六")
		} else if s == "u" {
			fmt.Println("星期天")
		} else {
			fmt.Println("输入错误")
		}
	} else {
		fmt.Println("输入错误")
	}
}

func main() {
	f1()
}
