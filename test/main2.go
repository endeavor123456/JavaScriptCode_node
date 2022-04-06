package main

import "fmt"

type Student3 struct {
	Age  int
	Name string
}
type Person2 struct {
	Name string
	Age  int
}
type Person3 struct {
	name string
	age  int
	sex  string
}
type Int int

func (p Person3) PrintInfo() {
	fmt.Println("姓名", p.name, "年龄", p.age)
}
func (i *Int) PPI() {
	fmt.Println(*i)
}
func main3() {
	var p = new(Person3)
	p.PrintInfo()
	var i Int
	i.PPI()
}
