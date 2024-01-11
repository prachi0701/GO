package main

import (
	"fmt"
)

func main() {
	// Array declaration and initialization
	var a [3]int
	a[0] = 12
	a[1] = 78
	a[2] = 50
	fmt.Println("Array a:", a)

	// Short hand declaration
	b := [3]int{12, 78, 50}
	fmt.Println("Array b:", b)

	// Partial initialization
	c := [3]int{12}
	fmt.Println("Partially initialized array c:", c)

	// Array with compiler-determined length
	d := [...]int{12, 78, 50}
	fmt.Println("Array d with compiler-determined length:", d)

	// Array value types
	e := [...]string{"USA", "China", "India", "Germany", "France"}
	f := e
	f[0] = "Singapore"
	fmt.Println("Original array e:", e)
	fmt.Println("Modified array f:", f)

	// Length of an array
	fmt.Println("Length of array e:", len(e))

	// Iterating arrays using range
	sum := 0.0
	g := [...]float64{67.7, 89.8, 21, 78}
	for i, v := range g {
		fmt.Printf("%dth element of g is %.2f\n", i, v)
		sum += v
	}
	fmt.Printf("Sum of all elements of g: %.2f\n", sum)

	// Multidimensional arrays
	var h [3][2]string
	h[0][0] = "apple"
	h[0][1] = "samsung"
	h[1][0] = "microsoft"
	h[1][1] = "google"
	h[2][0] = "AT&T"
	h[2][1] = "T-Mobile"
	fmt.Println("2D Array h:", h)

}
