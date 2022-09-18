package main

// 布尔值

import "fmt"

func main() {
	age := 19
	r := age >= 18
	if r {
		fmt.Println("你已经成年了")
	} else {
		fmt.Println("你还未成年")
	}

	count := 10
	for i := 0; i < count; i++ {
		fmt.Printf("i: %v\n", i)
	}
}
