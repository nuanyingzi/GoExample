package main

import "fmt"

// 指针

func main() {
	var ip *int
	fmt.Printf("ip: %v ip type: %T\n", ip, ip)
	var i int = 100
	ip = &i
	fmt.Printf("ip: %v ip type: %T\n", ip, ip)
	fmt.Printf("i: %v\n", *ip)

}
