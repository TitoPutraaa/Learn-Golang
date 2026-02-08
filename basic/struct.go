package main

import "fmt"

type Customer struct { // struct
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name ,"my name is", customer.Name)
}

type hasNae interface {
	GetName() string
	GetAge() int
}

func Nyapa(value hasNae) {
	fmt.Println("Haloo my name is", value.GetName(), "im", value.GetAge(), "years old")
}

type Person struct {
	name string
	age int
}

type Animal struct {
	name string
	age int
}

func (person Person) GetName() string {
	return person.name
}

func (person Person) GetAge() int {
	return person.age
}

func (animal Animal) GetName() string{
	return animal.name
}

func (animal Animal) GetAge() int {
	return animal.age
}


func main() {
	var tito Customer
	fmt.Println(tito)
	tito.Name = "Tito"
	tito.Address = "Jakarta"
	tito.Age = 21
	fmt.Println(tito)
	fmt.Println(tito.Name)

	putra := Customer {
		Name: "Putra",
		// Address: "Bali",
		Age: 22,
	}
	fmt.Println(putra)
	fmt.Println(putra.Name)

	bamulo := Customer{"Bamulo", "Bandung", 20}
	fmt.Println(bamulo)

	gojo := Customer{"Gojo", "Tokyo", 25}
	bamulo.sayHello("jupri")
	gojo.sayHello("asep")
	abdul := Person{name : "Abdul bin sulaiman", age : 90}
	Nyapa(abdul)
	jerry := Animal{"Jerry", 1}
	Nyapa(jerry)
	
	
}