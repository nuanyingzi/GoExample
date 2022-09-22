package main

import "fmt"

// 格式化输出

type WebSite struct {
	Name string
}

func main() {
	site := WebSite{
		Name: "duke",
	}
	fmt.Printf("site: %v\n", site)
	fmt.Printf("site: %#v\n", site) // 格式化输出 显示结构
	fmt.Printf("site: %T\n", site)

	a := 100
	fmt.Printf("a: %T\n", a)

	b := false
	fmt.Printf("b: %t\n", b)

	i := 98
	fmt.Printf("i: %v\n", i)    //
	fmt.Printf("i: %b\n", i)    // 二进制
	fmt.Printf("i: %c\n", i)    // 八进制
	fmt.Printf("i: %d\n", i)    // 十进制
	fmt.Printf("i: %x\n", 1234) // 十六进制

	x := 100
	p := &x
	fmt.Printf("p: %p\n", p) // 指针

}
