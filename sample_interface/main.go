package main

import "fmt"

type A interface {
	Do()
}

type B interface {
	Does()
}

type C struct {
	Name string
}

func (c *C) Do() {
	fmt.Println("Do: " + c.Name)
}

func (c *C) Does() {
	fmt.Println("Does: " + c.Name)
}

func hoge(c C) (A, B, C) {
	return &c, &c, c
}

func main() {
	c := C{Name: "fuga"}
	a, b, c := hoge(c)
	a.Do()
	b.Does()
	c.Do()
	c.Does()
	fmt.Println(c.Name)
}
