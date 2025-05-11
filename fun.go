package main

import "fmt"

func main() {
	fmt.Println("hi go")
	result := getResult(2, 5)
	fmt.Println("result : ", result)

	// function with multiple return type
	res, res1 := multipleReturn(4, 5, 4.0, 16.0)
	fmt.Println("res: ", res)
	fmt.Println("res1: ", res1)
}

// if the function name is start with small letter means
// it is not exportable
// but if it is capital letter then exportable
func getResult(a, b int) int {
	return a * b
}

func multipleReturn(a int, b int, c, d float64) (pro int, quo float64) {
	// pro = a * b
	// quo = c / d
	// return pro, quo
	return a * b, c / d
}
