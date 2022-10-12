package main

import "fmt"

// 构造方法

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) (*Person, error) {
	if name == "" {
		return nil, fmt.Errorf("name不能为空")
	}
	if age <= 0 {
		return nil, fmt.Errorf("age必须大于0")
	}
	return &Person{
		name,
		age,
	}, nil
}

func main() {
	person, error := NewPerson("dabao", 0)
	if error == nil {
		fmt.Printf("person: %v\n", *person)
	} else {
		fmt.Printf("err: %v\n", error)
	}
}
