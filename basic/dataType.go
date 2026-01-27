package main

import "fmt"

func main() {
	//overview
	fmt.Println("overview")
	thisArray := [...]string{"this", "make", "array"} // value fix
	// thisArray := [3]string{"this", "make", "array"}  
	thisSlice := []string{"this", "make", "Slice"} // can add value
	fmt.Println(thisArray)
	fmt.Println(thisSlice)


	// Aray
	fmt.Println("---Array---")
	var names [3]string

	names[0] = "tito"
	names[1] = "putra"
	names[2] = "bamulo"
	fmt.Println(names)

	var values = [3]byte {
		12, 53,
	}
	fmt.Println(values)
	
	fmt.Println(len(values))
	fmt.Println(values[2])
	values[2] = 99
	fmt.Println(values)

	var val = [...]byte {
		12, 53,
	}
	fmt.Println(val)

	// Slice 
	fmt.Println("---Slice---")
	slice := names[1:] // array[low : hight]
	slice2 := names[0:1]
	fmt.Println(slice)
	fmt.Println(slice2)
	
	fmt.Println("lengt of slice", len(slice))	
	fmt.Println("capacity of slice", cap(slice))	
	newSl := append(slice, "edward") // make new arr because capacity of the array in full
	fmt.Println("real array", names)
	fmt.Println("real slice",slice)
	fmt.Println("slice after append", newSl) 

	// modify slice
	makeSlice := []int { 12, 23, 43}
	// newSlice := make([]int16, 2, 3) //determine len and cap
	fmt.Println(makeSlice)
	fmt.Println(len(makeSlice))
	fmt.Println(cap(makeSlice))
	app := append(makeSlice, 55)
	fmt.Println(app)
	fmt.Println(len(app))
	fmt.Println(cap(app))
	app[0] = 32
	fmt.Println(app)

	// make new slice
	newSlice := make([]int16, 2, 3)
	fmt.Println(newSlice)
	newSlice[0] = 43
	newSlice[1] = 99
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	newSlice2 := append(newSlice, 43)
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))
	newSlice2[0] = 11 // still using same arr because enough capacity
	fmt.Println(newSlice)
	fmt.Println(newSlice2)
	newSlice3 := append(newSlice2, 100)
	fmt.Println(newSlice3)
	newSlice3[0] = 200 //using differecne arr because outof capacity
	fmt.Println(newSlice) //i[0] 11
	fmt.Println(newSlice3) // i[0] 200
    
	// copy slice
	fromSlice := names[:]
	toSlic := make([]string, 3, 5)
	copy(toSlic, fromSlice)
	fmt.Println(toSlic)
	fmt.Println(len(toSlic))
	fmt.Println(cap(toSlic))

	//Map
	fmt.Println("---MAP---")
	makeMap := make(map[string] string)
	fmt.Println(makeMap)
	newMap := map[string]string {
		"name" : "tito",
		"addr" : "bali",
		"sex" : "male",
	}
	fmt.Println(newMap)
	fmt.Println(newMap["name"])
	fmt.Println(newMap["addr"])
	fmt.Println("length", len(newMap))
	newMap["name"] = "titooo"
	fmt.Println(newMap)
	delete(newMap, "sex")
	fmt.Println(newMap)

}