package main

import "fmt"

func main() {
	name := "tito"
	// if statement
	fmt.Println("---If statement---")
	if name == "tito" {
		fmt.Println("Correct if")
	} else if name == "budi" {
		fmt.Println("correct else if ")
	} else {
		fmt.Println("correct else")
	}

	if nLength := len(name); nLength <= 5 {
		fmt.Println("name is short")
	}else {
		fmt.Println("long name")
	}

	fmt.Println("---Switch statement---")
	fmt.Println("---Loop statement---")
}