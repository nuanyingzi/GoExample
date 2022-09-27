package main

import "fmt"

func test2() {
	for i := 1; i <= 10; i++ {
	MYLABEL:

		for j := 1; j <= 10; j++ {

			if i == 2 && j == 2 {
				continue MYLABEL
			}
			fmt.Printf("%v %v\n", i, j)
		}
	}
}

func main() {
	test2()
}
