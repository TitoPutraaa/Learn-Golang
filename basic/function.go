package main

import (
	"fmt"
	"strconv"
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

func asValue(use string) string {
	return "this using " + use
}

func asParam(name string, position string, salary func(string) int) string {
	convString := strconv.Itoa(salary(position))
	result := name + " as " + position + " with salary " +  convString
	return result
}

func asParamFilter(position string) int {
	switch position {
		case "CEO" : 
			return 9000
		case "CTO" :
			return 7000
		default :
			return -1
	}
}

type paramFilter func(string)int

func asParam2(name string, position string, salary paramFilter) string {
	convString := strconv.Itoa(salary(position))
	result := name + " as " + position + " with salary " +  convString
	return result
}

type anon func(string) bool
func anonim(name string, blok anon) string{
	if  blok(name){
		return name + " you are blocked "
	} else {
		return "welcome " + name
	}
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
	funcToVar := asValue
	funcToVar2 := asValue
	fmt.Println(funcToVar("function as value"))
	fmt.Println(funcToVar2("function as value with difference var name"))
	res := asParam("tito", "CEO", asParamFilter)
	res2 := asParam2("putra", "CTO", asParamFilter)
	fmt.Println(res)
	fmt.Println(res2)
	blocked := func(name string) bool {return name == "tito"}
	fmt.Println(anonim("tito", blocked))
	fmt.Println(anonim("putra", func(name string) bool {return name== "tito"} ))
	

}