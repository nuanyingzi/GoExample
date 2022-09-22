package main

// if

import "fmt"

func main() {
	a, b := 1, 2
	if a > b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}

	if age := 20; age > 18 {
		fmt.Println("成年")
	} else {
		fmt.Println("未成年")
	}
}
