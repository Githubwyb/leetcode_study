package main

import "fmt"

type A struct {
	name string
}

// 实现String方法后，打印时会默认调用此方法
func (a A) String() string {
	return "A name: " + a.name
}

func main() {
	// 原型模式就是需要提供clone的接口，实现对对象的拷贝
	a := A{name: "a"}
	fmt.Println(&a) // A name: a 默认调用String方法
}
