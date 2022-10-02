package main

import "fmt"

// 指针数组

func main() {
	a := [3]int{1, 2, 3}
	var ap [3]*int
	fmt.Printf("ap: %v\n", ap)

	for i := 0; i < len(a); i++ {
		ap[i] = &a[i]
	}
	fmt.Printf("ap: %v\n", ap)
}
