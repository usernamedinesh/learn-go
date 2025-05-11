package main

import "fmt"

func main() {
	fmt.Println("variable")
	var name string = "dinesh"
	fmt.Println("name: ", name)

	var age int = 23
	fmt.Println("age:", age)

	// zero values
	var name1 string
	var age1 int
	var isTrue bool
	fmt.Println("name1: ", name1)
	res := (name1 == "")
	fmt.Println("is name1 empty: ", res) // true
	fmt.Println("age1:", age1)
	fmt.Println(isTrue) // false

	var (
		salary int    = 1000
		city   string = "chennai"
	)
	const (
		pi = 3.14
	)
	fmt.Println("salary: ", salary)
	fmt.Println("city: ", city)
	fmt.Println("pi: ", pi)
}
