package main

import "fmt"

func test1() {
	for i := 0; i < 10; i++ {
		if i == 9 {
			goto END
		}
		fmt.Printf("i: %v\n", i)
	}

END:
	fmt.Println("END")
}

func main() {
	test1()

}
