package main

import "fmt"

func test1() {
	m1 := map[string]string{
		"name":  "Tom",
		"age":   "18",
		"email": "852947475@qq.com",
	}
	for v, k := range m1 {
		fmt.Printf("v: %v, k: %v\n", v, k)
	}
}

func main() {
	test1()

}
