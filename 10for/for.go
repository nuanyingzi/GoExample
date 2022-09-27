package main

import "fmt"

func f1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// 初始条件可以写到外面
func f2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)

	}
}

// 结束条件可以写里面
func f3() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func main() {
	//f1()
	//f2()
	f3()
}
