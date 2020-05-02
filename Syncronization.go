package main

import (
	"fmt"
	"sync"
	"time"
)

func add(a *int, wait_group *sync.WaitGroup) {
	wait_group.Add(1) //Adding value to Wait Group, Counter Increases by 1
	for i := 0; i < 10; i++ {
		*a = *a + 1
		fmt.Println("Add() value  -", *a)
	}
	wait_group.Done() //Adding value to Wait Group, Counter Decreases by 1
}

func sub(a *int, wait_group *sync.WaitGroup) {
	wait_group.Add(1) //Adding value to Wait Group, Counter Increases by 1
	for i := 10; i > 0; i-- {
		*a = *a - 1
		fmt.Println("Sub() value  -", *a)
	}
	wait_group.Done() //Adding value to Wait Group, Counter Decreases by 1
}

func main() {
	start_time := time.Now()
	var a int = 10000
	var wait_group sync.WaitGroup // Synchronization in Gooooooooooooooooo.................
	go add(&a, &wait_group)
	go sub(&a, &wait_group)
	fmt.Println("Main Routine")
	time.Sleep(time.Second)
	end_time := time.Since(start_time)
	fmt.Println("Time taken is : %s", end_time)
	wait_group.Wait() // check wait group value == 0, then end the main routine
}
