package main

// defer
// 栈

import "fmt"

func main() {
	fmt.Println("Start")
	defer fmt.Println("Step1")
	defer fmt.Println("Step2")
	defer fmt.Println("Step3")
	fmt.Println("End")
}
