package main

import "fmt"

// defer : get execute after function complated, if that function error defer function still get execute
// panic : use for stop the program
func dMust() {
	fmt.Println("defer function get execute")
}

func runApp(error bool) {
	defer dMust() // if error true still get excute
	if error {
		panic("ups ERROR")
	}
	fmt.Println("running the Application")
}


func main() {
	runApp(true)
}