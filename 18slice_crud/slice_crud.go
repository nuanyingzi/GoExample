package main

import "fmt"

// 切片的增删改查

// add
func add() {
	var s1 []int
	s1 = append(s1, 100)
	s1 = append(s1, 200)
	fmt.Printf("s1: %#v\n", s1)
}

// del
func del() {
	var s1 = []int{1, 2, 3, 4, 5}
	s1 = append(s1[:2], s1[3:]...)
	fmt.Printf("s1: %#v\n", s1)
}

func copyTest() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 5)
	copy(s2, s1)
	s2[0] = 100
	fmt.Printf("s1: %#v\n", s1)
	fmt.Printf("s2: %#v\n", s2)
}

func main() {
	//add()
	//del()
	copyTest()
}
