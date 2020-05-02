package main

import (
	"fmt"
	"sort"
	"sync"
)

func sorting(arr []int, c chan []int, wg *sync.WaitGroup) {
	wg.Add(1)
	sort.Ints(arr)
	c <- arr
	//close(c)
	wg.Done()
}

func merge(a []int, b []int, c chan []int, wg *sync.WaitGroup) {
	wg.Add(1)
	res := make([]int, len(a)+len(b))
	i := 0
	j := 0
	k := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res[k] = a[i]
			i++
			k++
		} else {
			res[k] = b[j]
			j++
			k++
		}
	}
	for i < len(a) {
		res[k] = a[i]
		i++
		k++
	}
	for j < len(b) {
		res[k] = b[j]
		j++
		k++
	}

	c <- res
	wg.Done()
	//close(c)
}

func main() {
	l := 0
	var wg sync.WaitGroup

	fmt.Scanln(&l)
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		fmt.Scanln(&arr[i])
	}

	c := make(chan []int)
	parsize := l / 4
	go sorting(arr[:parsize], c, &wg)
	go sorting(arr[parsize:2*parsize], c, &wg)
	go sorting(arr[2*parsize:3*parsize], c, &wg)
	go sorting(arr[3*parsize:4*parsize], c, &wg)
	x1 := <-c
	x2 := <-c
	x3 := <-c
	x4 := <-c
	go merge(x1, x2, c, &wg)
	go merge(x3, x4, c, &wg)
	x5 := <-c
	x6 := <-c
	go merge(x5, x6, c, &wg)
	wg.Wait()
	x7 := <-c
	fmt.Println("Sorted Results are -")
	for i := range x7 {
		fmt.Printf("%d \t", x7[i])
		//fmt.Println(x[i])
	}

}
