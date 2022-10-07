package main

import "fmt"

type USBER interface {
	read()
	write()
}

type Computer struct {
	name string
}

func (c Computer) read() {
	fmt.Printf("c.name: %v\n", c.name)
	fmt.Println("read...")
}

func (c Computer) write() {
	fmt.Printf("c.name: %v\n", c.name)
	fmt.Println("write...")
}

type Mobile struct {
	name string
}

func (m Mobile) read() {
	fmt.Printf("m.name: %v\n", m.name)
	fmt.Println("mobile read...")
}

func (m Mobile) write() {
	fmt.Printf("m.name: %v\n", m.name)
	fmt.Println("mobile write...")
}

func main() {
	c := Computer{
		name: "微星",
	}
	c.read()
	c.write()

	m := Mobile{name: "诺基亚"}
	m.read()
	m.write()
}
