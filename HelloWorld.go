package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int { //function syntax  - func <name> (varaible type) retuntype{ //code }
	res := x + y
	return res
}

func cal(x, y int) (int, int) {
	res := x + y
	sub := x - y
	return res, sub
}

func cal1(x, y int) (res, sub int) { // function return variables can be delcared in return type only
	res = x + y
	sub = x - y
	return // this will return "res and Sub" variables // same as cal func, but diff syntax
}

func main() {
	fmt.Println("Hello World")

	var result int = 20

	//i := 1 //variable devalaration -- var i int = 1]
	for i := 0; i < 5; i++ {
		fmt.Println("for loop works -", i)
		//i++
	}
	//functions
	var f float64 = 10
	x := 20
	y := 10
	result = add(x, y) //function
	fmt.Println(result)
	res, sub := cal(x, y)
	//fmt.Println(res, " ", sub)
	fmt.Println(res, " ", sub)
	var floatres = math.Sqrt(f)
	var intres = int(math.Ceil(floatres))
	fmt.Printf("Rounding to - %.2f  \n", floatres) //placeholder for float is - g(or)f is used
	fmt.Printf("Rounding to - %d", intres)
}
