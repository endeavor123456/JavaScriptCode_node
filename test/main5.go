package main

import "fmt"

type test5 interface {
	demo3(string, int) int
}
type demo2 struct {
	name string
	age  int
}
type demo4 struct {
	demo2
	sex string
}

func (d *demo2) demo3(a string, b int) (sum int) {
	d.name = a
	d.age = b
	sum = d.age
	return

}
func main05() {
	var d = demo2{
		name: "狗子",
		age:  20,
	}
	d2 := demo2{
		name: "狗子",
		age:  20,
	}
	var d3 = &demo2{
		name: "狗子",
		age:  20,
	}
	fmt.Println(d.name)
	fmt.Println(d2)
	fmt.Println(d3.age)
	// d2.demo3("狗子1号", 20)
	fmt.Println(d)

	var d5 = demo4{
		demo2{
			age:  10,
			name: "狗子3号",
		},
		"男", //注意下这个地方  前面的字段名没写上后边的也不要写  要协调
	}
	fmt.Println(d5)
	var a test5
	a = d3
	fmt.Println(a)
	d2.demo3("狗子3号", 20)
	//接口里的方法限制多一些
}
