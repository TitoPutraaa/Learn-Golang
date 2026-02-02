package main

import "fmt"

type Costemer struct {
	Name, Address string
	Age           int
}

func main() {
	var tito Costemer
	tito.Name = "Tito"
	tito.Address = "Jakarta"
	tito.Age = 21

	fmt.Println(tito)
	fmt.Println(tito.Name)
}