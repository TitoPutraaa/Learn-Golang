package main

import "fmt"

func main () {
	// basic
	var name string
	name = "tito"
	fmt.Println(name)
	fmt.Println("---------------")
	
	var midName = "putra"
	fmt.Println(midName)
	fmt.Println("---------------")

	lastName := "Bamulo"
	fmt.Println(lastName)
	lastName = "Bamulo00" // change value must without ":="
	fmt.Println(lastName)

	// multipele variable
	var (
		age = 21
		sex = "male"
	)
	fmt.Println("age of", age, ", sex", sex)

	const (
		hobby = "basket ball"
		favFood = "sate"
	)
	fmt.Println("hobby", hobby, ", Favorite food" , favFood)

	//convertion data type

	var num32 int32 = 1323
	var num64 int64 = int64(num32)

	fmt.Println(num32)
	fmt.Println(num64)

	var me = "tito"
	var e = me[1]
	fmt.Println("byte char", e)
	var eString = string(e)
	fmt.Println("covert byte to string", eString)

	// type declaration 
	type nim string //make new alias
	var nimTito nim = "12314123"
	fmt.Println(nimTito)
	fmt.Println(nim("53453"))
}
	
