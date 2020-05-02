package main

import (
	"fmt"
	"sort"
	"strconv"
)

// func sorting(arr []int) {
// 	sort.Ints(*arr)
// }

func main() {
	arr := make([]int, 0, 3)
	for i := 0; i < 3; i++ {
		var x string
		fmt.Scanln(&x)
		if x == "X" {
			//os.Exit(1)
			break
		} else {
			k, err := strconv.Atoi(x)
			if err == nil {
				arr = append(arr, k)
				sort.Ints(arr)
				//sorting(&arr)
			} else {
				fmt.Println("Invalid int -", x)
			}

		}

	}

	for i := range arr {
		fmt.Printf("%d \t", arr[i])
	}
}
