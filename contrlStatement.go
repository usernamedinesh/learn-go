package main

import "fmt"

func main() {
	fmt.Println("hello go")

	age := 23

	// if
	if age >= 23 {
		fmt.Println("you are an adult!")
	} else if age > 18 {
		fmt.Println("you are an minor!")
	} else {
		fmt.Println("enter an valid age!")
	}

	// switch
	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("start of the week")
	case "Tuesday", "WebnessDay":
		fmt.Println("middle of week")
	default:
		{
			fmt.Println("defult is hitting")
		}
	}

	// loop
	// only have for loop
	for i := 0; i < 3; i++ {
		fmt.Println("nums: ", i)
	}

	// use this instead of while
	counter := 0
	for counter < 3 {
		fmt.Println("counter: ", counter)
		counter++
	}

	// inifinite loop
	count := 0
	for {
		if count == 3 {
			break
		}
		fmt.Println("count: ", count)
		count++
	}
}
