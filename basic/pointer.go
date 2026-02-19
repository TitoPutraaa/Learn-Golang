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
	variable2.City = "Bogor"

	fmt.Println(variable1) // still first value
	fmt.Println(variable2) // make copy of first value

	// USING PASS BY REFFERENCE
	var variabel3 *Address = &variable1
	variabel3.City = "Depok"
	fmt.Println(variable1) //variable 1 is change becuase of var3
	fmt.Println(variabel3)

	// ASTERISK OPERATOR
	variabel3 = &Address{"Surabaya", "Jawa Tengah", "Indonesia"}
	fmt.Println(variabel3) // change var 3 to point other data
}
