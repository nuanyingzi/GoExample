package main

// 切片

import "fmt"

func test1() {
	var s1 []int
	var s2 []string
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
}
func test2() {
	var s1 = make([]int, 2)
	fmt.Printf("s1: %v\n", s1)
}

func test3() {
	s1 := []int{1, 2, 3}
	fmt.Printf("len(s1): %v, cap(s1): %v\n", len(s1), cap(s1))

}

func main() {
	//test1()
	//test2()
	test3()
}
