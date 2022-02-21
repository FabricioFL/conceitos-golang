package main

import "fmt"

func main() {

	a := 10
	b := a << 1
	fmt.Printf("%d %b %#x\n%d %b %#x", a, a, a, b, b, b)
}
