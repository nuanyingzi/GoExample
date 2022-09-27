package main

import "fmt"

func f1() {
	var a = [...]int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("v: %v\n", v)
	}
}

func f2() {
	var s = []int{12, 23, 34, 45}
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	//f1()
	f2()
}
