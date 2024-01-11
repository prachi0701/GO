package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Boolean Type
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)

	// Signed Integers
	var signedInt int = 89
	fmt.Println("value of signedInt is", signedInt)
	fmt.Printf("type of signedInt is %T, size of signedInt is %d\n", signedInt, unsafe.Sizeof(signedInt))

	// Unsigned Integers
	var unsignedInt uint = 255
	fmt.Println("value of unsignedInt is", unsignedInt)
	fmt.Printf("type of unsignedInt is %T, size of unsignedInt is %d\n", unsignedInt, unsafe.Sizeof(unsignedInt))

	// Floating Point Types
	f1, f2 := 5.67, 8.97
	fmt.Printf("type of f1 %T f2 %T\n", f1, f2)
	fmt.Println("sum", f1+f2, "diff", f1-f2)

	// Complex Types
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	cmul := c1 * c2
	fmt.Println("sum:", cadd)
	fmt.Println("product:", cmul)

	// String Type
	first := "Prachi"
	last := "Bhati"
	name := first + " " + last
	fmt.Println("My name is", name)

	// Type Conversion
	i := 55           // int
	j := 67.8         // float64
	sum := i + int(j) // j is converted to int
	fmt.Println("sum", sum)
}
