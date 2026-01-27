package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("halooo")
}

func sayHelloTo(fName string, lName string) {
	fmt.Println("haloo", fName, lName)
}

func reHello(fName string) string {
	return "halo " + fName
}

func person() (string, uint)  { //retun multiple value 
	return "Tito", 21
}

func person2(name string) (fName, lName string, age int) {
	fName = "tito"
	lName = "bamulo"
	age = 21
	fName = name

	return fName, lName, age
}

func sumAll(nums ...int) int { //variable argument, add without make slice
	result := 0

	for _, count := range nums {
		result += count
	}
	return result
}

func sumAll2(nums []int) int { //add make slice first
	result := 0

	for _, count := range nums {
		result += count
	}
	return result
}

func sumAll3(nums ...int) int { 
	result := 0

	for _, count := range nums {
		result += count
	}
	return result
}



func main() {
	sayHello()
	sayHelloTo("tito", "putra")
	sayHelloTo("joko", "widodo")
	result := reHello("bamulo")
	fmt.Println(result)
	name, age := person()
	fmt.Println("name", name, "\nage ", age)
	a, b, c := person2("mike")
	fmt.Println(a, b, c)
	fmt.Println(sumAll(10,10,10)) // without add new slice 
	data := []int {10,10,10,5}
	fmt.Println(sumAll2(data)) // add new slice 
	fmt.Println(sumAll3(data...)) // using variable argument, when already have slice
}