package main

import "fmt"

const (
	_ = iota + 2022
	a = iota + 2022
	b = iota + 2022
	c = iota + 2022
	d = iota + 2022
)

func main() {

	fmt.Printf(" %v \n %v \n %v \n %v", a, b, c, d)
}
