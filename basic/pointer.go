package main

import "fmt"

// go default is using pass by value

type Address struct {
	City, Province, Country string
}

func main() {
	// USING PASS BY VALUE
	variable1 := Address{"Banten", "Jawa Barat", "Indonesia"}
	var variable2 Address = variable1
	variable2.City = "Bandung"

	fmt.Println(variable1) // still first value
	fmt.Println(variable2) // make copy of first value

	// USING PASS BY REFFERENCE
	// var variabel3 *Address = &variable1
	// variabel3.City = "Depok"
	// fmt.Println(variable1)
	// fmt.Println(variabel3)
}
