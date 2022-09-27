package main

import "fmt"

// 判断map中的一个key是否存在
func test1() {
	m := map[string]string{
		"name":  "Tom",
		"age":   "18",
		"email": "852947475@qq.com",
	}
	k1 := "age"
	v1, ok := m[k1]
	fmt.Printf("v1: %v, ok: %v\n", v1, ok)
	k2 := "age1"
	v2, ok := m[k2]
	fmt.Printf("v2: %v, ok: %v\n", v2, ok)

}

func main() {
	var m1 = make(map[string]string)
	m1["name"] = "Tom"
	m1["age"] = "18"
	m1["email"] = "852947475@qq.com"
	fmt.Printf("m1: %v\n", m1)

	m2 := map[string]string{
		"name":  "Tom",
		"age":   "18",
		"email": "852947475@qq.com",
	}
	fmt.Printf("m2: %v\n", m2)

	test1()

}
