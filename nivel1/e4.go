package main

import "fmt"

type numero int64

var x numero

func main() {

	fmt.Printf("\n %T %v", x, x)
	x = 42
	fmt.Printf("\n %T %v", x, x)
}
