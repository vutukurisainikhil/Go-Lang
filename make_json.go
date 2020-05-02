package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {

	var nameperson string
	var addr string
	//p := new(person)
	//var p person
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter an name")
	scanner.Scan()
	nameperson = scanner.Text()
	fmt.Println("Enter an address")
	scanner.Scan()
	addr = scanner.Text()

	// var userInfo map[string]string
	// userInfo = make(map[string]string)

	// userInfo["name"] = name
	// userInfo["address"] = addr
	// p.name = nameperson
	// p.address = addr
	p := person{Name: nameperson, Address: addr}
	fmt.Printf("%s\n", p)
	jsonUserInfo, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json Marshaling failed")
		return
	}
	fmt.Printf("%s\n", string(jsonUserInfo))

	//p1 := new(person)
	var p1 person
	erro := json.Unmarshal(jsonUserInfo, &p1)
	if erro == nil {
		fmt.Println(p1.Name)
		fmt.Println(p1.Address)
	}
	fmt.Println(p1.Name)
	fmt.Println(p1.Address)
}
