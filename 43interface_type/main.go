package main

import "fmt"

// 一个类型可以实现多个接口

type Player interface {
	playMusic()
}

type Video interface {
	playVideo()
}

type Mobile struct {
	name string
}

func (m Mobile) playMusic() {
	fmt.Println("播放音乐")
}

func (m Mobile) playVedio() {
	fmt.Println("播放视频")
}

// 多个类型实现一个接口

type Pet interface {
	eat()
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func (d Dog) eat() {
	fmt.Printf("%v dog eat...\n", d.name)
}

func (c Cat) eat() {
	fmt.Printf("%v cat eat...\n", c.name)
}

func main() {
	// 一个类型可以实现多个接口
	m := Mobile{name: "小米"}
	m.playMusic()
	m.playVedio()

	// 多个类型实现一个接口
	/*d := Dog{name: "大黄"}
	d.eat()
	c := Cat{name: "喵喵"}
	c.eat()*/
	var p Pet
	p = Dog{name: "大黄"}
	p.eat()
	p = Cat{name: "喵喵"}
	p.eat()

}
