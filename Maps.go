package main

import (
	"fmt"
)

type employee struct {
	salary  int
	country string
}

func main() {
	// Creating and initializing a map
	employeeSalary := make(map[string]int)
	employeeSalary["steve"] = 12000
	employeeSalary["jamie"] = 15000
	employeeSalary["mike"] = 9000

	// Printing map
	fmt.Println("employeeSalary map contents:", employeeSalary)

	// Checking if a key exists and retrieving value
	employee := "jamie"
	salary, exists := employeeSalary[employee]
	if exists {
		fmt.Println("Salary of", employee, "is", salary)
	} else {
		fmt.Println(employee, "not found")
	}

	// Iterating over a map
	fmt.Println("Contents of the map:")
	for key, value := range employeeSalary {
		fmt.Printf("employeeSalary[%s] = %d\n", key, value)
	}

	// Deleting an item
	delete(employeeSalary, "steve")
	fmt.Println("map after deletion", employeeSalary)

	// Length of the map
	fmt.Println("length of employeeSalary is", len(employeeSalary))

}

