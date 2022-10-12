package main

import (
	"fmt"
	"time"
)

// 并发

func show(msg string) {
	for i := 1; i < 6; i++ {
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go show("PHP") // 启动了一个协程来执行
	go show("Golang")
	time.Sleep(time.Millisecond * 2000)

	fmt.Println("main...")

}
