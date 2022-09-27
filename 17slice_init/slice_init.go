package main

import "fmt"

// 切片初始化

func test1() {
	s := []int{1, 2, 3, 4, 5}
	s1 := s[:3]
	s2 := s[3:]
	s3 := s[:]
	fmt.Printf("s1: %#v\n", s1)
	fmt.Printf("s2: %#v\n", s2)
	fmt.Printf("s3: %#v\n", s3)
}

func main() {
	test1()
}
