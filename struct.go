package main

import (
	"fmt"
)

// struct stand for structure
// They are data types that can hold data and be passed around the application.
// custom data type :=>  we can use this as an argument in function
// and can use as a  return type
type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("hi struct")
	person := Person{
		Name: "dinesh",
		Age:  23,
	}
	fmt.Printf("Person: %v\n", person)
	fmt.Printf("+Person: %+v\n", person)

	// anonymouns struct
	employee := struct {
		// strucs initialization
		salary  int
		geneder string
	}{
		// adding values
		salary:  400,
		geneder: "male",
	}
	fmt.Println("EMPLOYEE: ", employee)

	// nested struct
	type Address struct {
		City string
		Pin  int
	}

	type Contact struct {
		phone   int
		Address Address // it is an struct
	}
	conatact := Contact{
		phone: 888888,
		Address: Address{
			City: "tamulpur",
			Pin:  333,
		},
	}
	fmt.Println("nested contact: ", conatact)
	fmt.Println("before", person.Name)
	modifyPerson(person)
	// modifyPerson(&person) // if we pass address then it will actual modify the original value
	fmt.Println("after", person.Name)

	// and we can do like this
	x := 20
	ptr := &x // store the address of x
	*ptr = 30 // store the value in the address

	person.methodOnPerson("yes buddy")

	fmt.Println("modifiy the original value", person.Name)
}

// only goes copy here
// func modifyPerson(person *Person) { // if we pass address then receive like this
func modifyPerson(person Person) {
	person.Name = "temp name"
	fmt.Println("person inside the func : ", person.Name)
}

// lets add an method on type struct
// remember only exist on struct
func (p *Person) methodOnPerson(name string) {
	p.Name = name
	fmt.Println("name is : ", p.Name)
}
