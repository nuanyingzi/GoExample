package main

import "fmt"

// 接口的嵌套

type Flyer interface {
	fly()
}

type Swimmer interface {
	swim()
}

type FlyFish interface {
	Flyer
	Swimmer
}

type Fish struct {
}

func (f Fish) fly() {
	fmt.Println("Fish can fly")
}

func (f Fish) swim() {
	fmt.Println("Fish can swim")
}

func main() {
	var fs FlyFish
	fs = Fish{}
	fs.fly()
	fs.swim()
}
