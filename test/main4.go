package main

import "fmt"

type Test struct {
	name string
}

func test(t Test) {
	fmt.Println(t)
}
func main09() {
	// test(Test{"狗子"})
	x := 1
	{
		x := 2
		fmt.Print(x)
	}
	fmt.Print(x)
}
