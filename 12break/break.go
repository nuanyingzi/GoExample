package main

import "fmt"

func test1() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("i:%v\n", i)
		if i == 6 {
			break
		}
	}
}

func test2() {
	i := 2
	switch i {
	case 1:
		fmt.Println("A")
		break
	case 2:
		fmt.Println("B")
		fallthrough
	case 3:
		fmt.Println("C")
		break
	default:
		fmt.Println("NOTHING")

	}
}

func test3() {
MYLABEL:
	for i := 1; i <= 10; i++ {
		fmt.Printf("i:%v\n", i)
		if i >= 5 {
			break MYLABEL
		}
	}
	fmt.Println("--END--")
}

func main() {
	//test1()
	test2()
	//test3()
}
