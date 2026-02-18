package main

import (
	"fmt"
)

func random() interface{} {
	return "OK"
}

func main() {
	var value any = random()
	// var convValue string = value.(string)
	// fmt.Println(convValue)

	switch cVal := value.(type) {
	case string:
		fmt.Println("string", cVal)
	case int:
		fmt.Println("int", cVal)
	default:
		fmt.Println("Unknown", cVal)
	}

}
