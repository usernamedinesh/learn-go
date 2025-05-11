package main

import "fmt"

func main() {
	fmt.Println("hello array")

	nums := [4]int{1, 23, 43, 5}
	fmt.Printf("nums: %v\n", nums)

	fmt.Println("last value", nums[len(nums)-1]) // 5

	numsberInti := [...]int{1, 2, 3, 4, 5, 3}
	fmt.Printf("numsberInti: %v\n", numsberInti)

	maxtrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("maxtrix: %v\n", maxtrix)

	// slice (dynamic array)
	// we can make like slice portion of an array
	// like this

	// portionSlice := nums[:] // this become portion of an array here we have methods
	// declaring slice

	fruits := []string{"apple", "banana", "mango"}
	fmt.Printf("fruits: %v\n", fruits)
	fruits = append(fruits, "papaya") // adding into fruits we can add more
	fmt.Printf("fruits: %v\n", fruits)

	// more append
	moreFruits := []string{"one", "two"}
	fruits = append(fruits, moreFruits...)
	fmt.Printf("fruits: %v\n", fruits)

	// iteration
	for index, value := range nums {
		fmt.Printf("index %d, value: %d\n", index, value)
	}

	// here i want only value so i need to igore index
	for _, value := range nums {
		fmt.Printf(" value: %d\n", value)
	}

	// map:  key value
	capitalCities := map[string]string{
		"USA":   " washingTokn DB",
		"INDIA": "new delhi",
	}
	fmt.Println(capitalCities["USA"]) // washingTokn DB
	// chech exist or not
	capital, exist := capitalCities["INDIA"]
	if exist {
		fmt.Println(exist) // return true
		fmt.Println("this is the capital", capital)
	} else {
		fmt.Println("does not exsit")
	}
	// delete
	delete(capitalCities, "USA")
	fmt.Println(capitalCities)
}
