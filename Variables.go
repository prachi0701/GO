package main

import (
	"fmt"
	"math"
)

func main() {
	// Variable declaration and initialization
	var age int
	fmt.Println("My age is", age) 

	// Assigning values to a variable
	age = 22
	fmt.Println("My age is", age) 

	// Declaring a variable with an initial value
	var ageWithInitialValue int = 22
	fmt.Println("My age is", ageWithInitialValue) 

	// Short hand declaration
	count := 10
	fmt.Println("Count =", count) 

	// Short hand declaration for multiple variables
	nameShort, ageShort := "Prachi", 22
	fmt.Println("my name is", nameShort, "age is", ageShort) 

	// Runtime value assignment
	aRuntime, bRuntime := 145.8, 543.8
	cRuntime := math.Min(aRuntime, bRuntime)
	fmt.Println("Minimum value is", cRuntime) 
}
