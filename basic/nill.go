package main

import "fmt"

func newMap(name string) map[string]string { // nil can't do in primitive datatype
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := newMap("Roberto")
	if data == nil {
		fmt.Println("this map is nil or empyt")
	} else {
		fmt.Println("this map not nil")
	}
	fmt.Println(data)
}