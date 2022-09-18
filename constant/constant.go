package main

// 常量

import "fmt"

func main() {
	const PI float64 = 3.1415926
	fmt.Println(PI)

	const (
		WEIGHT = 200
		HEIGHT = 100
	)
	fmt.Printf("WEIGHT: %v\n", WEIGHT)
	fmt.Printf("HEIGHT: %v\n", HEIGHT)

	const (
		a1 = iota
		a2 = iota
		a3 = iota
	)

	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
}
