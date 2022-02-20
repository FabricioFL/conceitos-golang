package main

import "fmt"

type numero int64

var x numero
var y int

func main() {

	fmt.Printf("\n %T %v", x, x)
	x = 42
	fmt.Printf("\n %T %v", x, x)
	y = int(x)
	fmt.Printf("\n %T %v", y, y)
}
