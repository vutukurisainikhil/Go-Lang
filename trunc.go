package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Scanln(&x)
	fmt.Println(math.Trunc(x))
}
