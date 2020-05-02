package main

import (
	"fmt"
)

func add(a int) int {
	x := a + 1
	return x
}

func sub(a int) int {
	x := a + 1
	return x
}

func main() {
	a := 10
	b := add(a)
	fmt.Println("the value in add 1st func -", b)
	c := sub(a)
	fmt.Println("the value in add 1st func -", c)
}
