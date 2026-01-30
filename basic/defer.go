package main

import "fmt"

// defer : get execute after function complated, if that function error defer function still get execute
func dMust() {
	fmt.Println("defer function get execute")
}

func runApp() {
	defer dMust()
	fmt.Println("running the Application")
}

func main() {
	runApp()
}