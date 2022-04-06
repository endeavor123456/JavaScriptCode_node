package main

import "fmt"

type res struct {
	age int
}

func result() (sum res) {
	return
}
func main() {
	res := result()
	fmt.Println(res)

}
