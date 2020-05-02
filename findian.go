package main

import (
	"bufio"
	"fmt"
	"os"
	m "strings"
)

func main() {
	flag := false
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	line := scanner.Text()
	fmt.Println("captured:", line)
	//l := len(line)
	if m.Contains(line, "a") || m.Contains(line, "A") {
		if line[0] == 'i' || line[0] == 'I' {
			if line[len(line)-1] == 'n' || line[len(line)-1] == 'N' {
				flag = true
				fmt.Println("Found!")
			}
		}
	}
	if !flag {
		fmt.Println("Not Found!")
	}
}
