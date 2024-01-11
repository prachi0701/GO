package main

import (
	"fmt"
	"math"
)

func main() {
	// Declaring a Single Constant
	const a = 50
	fmt.Println("Single Constant:", a)

	// Declaring a Group of Constants
	const (
		name    = "Prachi"
		age     = 22
		country = "India"
	)
	fmt.Println("Group of Constants:", name, age, country)

	// String Constants, Typed and Untyped
	const hello = "Hello World" // Untyped constant
	var nameString string = hello
	fmt.Println("String Constant:", nameString)

	// Typed constant
	const typedHello string = "Typed Hello World"
	fmt.Println("Typed String Constant:", typedHello)

	// Numeric Constants
	const num = 5
	var intVar int = num
	var int32Var int32 = num
	var float64Var float64 = num
	var complex64Var complex64 = num
	fmt.Println("Numeric Constants:", intVar, int32Var, float64Var, complex64Var)

	// Numeric Expressions
	var numericExpression = 5.9 / 8
	fmt.Printf("Numeric Expression: type %T value %v\n", numericExpression, numericExpression)

	// Boolean Constants
	const trueConst = true
	type myBool bool
	var defaultBool = trueConst // allowed
	var customBool myBool = trueConst // allowed
	// defaultBool = customBool // not allowed due to strict type
	fmt.Println("Boolean Constants:", defaultBool, customBool)

	// Strong Typing in Go
	type myString string
	var defaultName string = "Bhati" // allowed
	var customName myString = "Bhati" // allowed
	fmt.Println("Strong Typing Example:", defaultName, customName)
}

