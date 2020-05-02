package main

import "fmt"

func swap(arr []int, j int) {
	arr[j+1], arr[j] = arr[j], arr[j+1]
}

func Bubblesort(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j)
			}
		}
	}
}

func main() {
	arr := make([]int, 0, 10)
	fmt.Println("Enter 10 Digits")
	for i := 0; i < 10; i++ {
		var k int
		fmt.Scanln(&k)
		arr = append(arr, k)
	}
	Bubblesort(arr)
	fmt.Println("The Sorted array -", arr)
}
