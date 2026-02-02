package main

import "fmt"

// defer : get execute after function complated, if that function error defer function still get execute
// panic : use for stop the program
// recover : to show error message
func dMust() {
	fmt.Println("defer function get execute")
	message := recover()
	fmt.Println(message, "error recover")
}

func runApp(error bool) {
	defer dMust() // if error true still get excute
	if error {
		panic("ups ERROR")
	}
	fmt.Println("running the Application")
}


func main() {
	runApp(false)
}