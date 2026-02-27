package main

import "fmt"

// go default is using pass by value

type Address struct {
	City, Province, Country string
}

func main() {
	// USING PASS BY VALUE default
	variable1 := Address{"Banten", "Jawa Barat", "Indonesia"}
	var variable2 Address = variable1
	variable2.City = "Bogor"

	fmt.Println(variable1, "by value, Var 1")       // still first value
	fmt.Println(variable2, "by value var 2 change") // make copy of first value
	fmt.Println()

	// USING PASS BY REFFERENCE
	var variabel3 *Address = &variable1
	fmt.Println(variable1)
	variabel3.City = "Depok"                                               // it's change parents variable, because var 3 is point to parent so when it's change that will change parents to
	fmt.Println(variable1, "by reference, var 1 get change because var 3") //variable 1 is change becuase of var3
	fmt.Println(variabel3, "by refference, var 3 change main value")
	fmt.Println(variabel3)

	// ASTERISK OPERATOR
	variabel3 = &Address{"Surabaya", "Jawa Tengah", "Indonesia"} // make new data
	fmt.Println()
	fmt.Println(variable1, "still reference first value")
	fmt.Println(variabel3, "make new value and reference to them") // var 3 is change because they makes new data

	*variabel3 = Address{"Banyuwangi", "Jawa Timur", "Indonesia"}
	fmt.Println(variable1)
	fmt.Println(variabel3)

	// PASS BY REFFERECE pointer
	fmt.Println("-------------BY Reference-----------")
	var num1 int = 67
	fmt.Println(&num1)
	fmt.Println(num1)
	var num2 *int = &num1  // num2 = addres of num1
	var num3 **int = &num2 // pointer nya pointe wkwkw
	// num2 := &num1
	fmt.Println(num2)
	fmt.Println(*num2)
	fmt.Println(num3)
	fmt.Println(*num3) // value of the addr
	*num2 = 1          // because num2 is contain addres so we have to add * so that will cange is the value in that addr not the addr
	// num2 = 3 ERROR because you can't change addr
	fmt.Println(num1)
	fmt.Println(*num2)
	// & = address of something, * = derefference / when you have addr it's will make you go to see the value in that addr

	// PASS BY VALUE default
	fmt.Println("-------------BY Value-----------")
	var ang1 int = 99
	fmt.Println(ang1)
	var ang2 int = ang1
	fmt.Println(ang2)
	ang2 = 666
	fmt.Println(ang1) // still 99
	fmt.Println(ang2) // make copy to 666
	fmt.Println(&ang1)
}
