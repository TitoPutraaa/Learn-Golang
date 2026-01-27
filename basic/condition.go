package main

import "fmt"

func main() {
	name := "tito"
	fmt.Println(name)
	// if statement
	fmt.Println("---If statement---")
	if name == "tito" {
		fmt.Println("Correct if")
	} else if name == "budi" {
		fmt.Println("correct else if ")
	} else {
		fmt.Println("correct else")
	}

	if nLength := len(name); nLength <= 5 { //short
		fmt.Println("name is short")
	}else {
		fmt.Println("long name")
	}

	// switch
	fmt.Println("---Switch statement---")
	switch name {
	case "tito":
		fmt.Println("the name tito")
	case "budi": 
		fmt.Println("the name budi")
	default : 
		fmt.Println("unknown name")
	} 


	switch cnLen := len(name); cnLen <= 5 {
	case true : 
			fmt.Println("short name")
	case false : 
		fmt.Println("long name")
	}

	// fmt.Println("---Loop statement---")
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("index ", i)
	// }

	// for i; name 
}