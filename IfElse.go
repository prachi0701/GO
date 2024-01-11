package main

import (
	"fmt"
)

func main() {
	// Example 1: Simple if statement to check if a number is even or odd
	num1 := 10
	if num1%2 == 0 {
		fmt.Println("The number", num1, "is even")
	} else {
		fmt.Println("The number", num1, "is odd")
	}

	// Example 2: if-else-if statement
	num2 := 99
	if num2 <= 50 {
		fmt.Println(num2, "is less than or equal to 50")
	} else if num2 >= 51 && num2 <= 100 {
		fmt.Println(num2, "is between 51 and 100")
	} else {
		fmt.Println(num2, "is greater than 100")
	}

	// Example 3: if with assignment
	if num3 := 10; num3%2 == 0 {
		fmt.Println(num3, "is even")
	} else {
		fmt.Println(num3, "is odd")
	}

	// Idiomatic Go - avoiding unnecessary else
	num4 := 10
	if num4%2 == 0 {
		fmt.Println(num4, "is even")
		return
	}
	fmt.Println(num4, "is odd")
}
