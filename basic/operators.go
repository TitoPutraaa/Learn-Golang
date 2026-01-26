package main

import "fmt"

func main() {
	// basic aritmatics 
	fmt.Println("---Aritmatic---")
	num := 1 + 3 - 1 * 2 / 2
	fmt.Println(num) 

	num += 2 // num = num + 2
	fmt.Println(num) 
	
	num++ 
	fmt.Println(num) 

	//comparation
	fmt.Println("---Comparation---")
	prsn1 := "tito"
	prsn2 := "tito"

	isSame := prsn1 == prsn2
	fmt.Println(isSame) // ... etc
	
}