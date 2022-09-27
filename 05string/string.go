package main

import (
	"bytes"
	"fmt"
	"strings"
)

// 字符串

func main() {
	// 双引号和反引号
	s := "Hello world"
	s1 := `
	line1
	line2
	line3
	`
	fmt.Printf("s: %v\n", s)
	fmt.Printf("s1: %v\n", s1)

	// 字符串连接1
	s2 := "Hello "
	s3 := "world"
	msg := s2 + s3
	fmt.Printf("msg: %v\n", msg)
	// 字符串连接2
	s4 := "tom"
	s5 := "20"
	msg1 := fmt.Sprintf("%s, %s", s4, s5)
	fmt.Printf("msg1: %v\n", msg1)
	// 字符串连接3
	name := "tom"
	age := "20"
	content := strings.Join([]string{name, age}, ",")
	fmt.Println(content)
	// 字符串连接4
	var buffer bytes.Buffer
	buffer.WriteString("Tom")
	buffer.WriteString(",")
	buffer.WriteString("20")
	fmt.Printf("buffer.String(): %v\n", buffer.String())

	// 字符串函数
	ss := "hello world"
	fmt.Printf("len(ss): %v\n", len(ss))
	fmt.Printf("strings.Split(s, \"\"): %v, type: %T\n", strings.Split(s, " "), strings.Split(s, " "))
	fmt.Printf("strings.Contains(s. \"hello\"): %v\n", strings.Contains(s, "hello"))
}
