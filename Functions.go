package main

import (
	"fmt"
)

// calculateBill function
func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

// rectProps function with multiple return values
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

// rectProps function with named return values
func rectPropsNamed(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return // No explicit return value needed
}

func main() {
	// Using calculateBill
	price, no := 90, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)

	// Using rectProps
	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f\n", area, perimeter)

	// Using rectProps with named return values
	areaNamed, perimeterNamed := rectPropsNamed(10.8, 5.6)
	fmt.Printf("Area  %f Perimeter  %f\n", areaNamed, perimeterNamed)

	// Using blank identifier
	areaOnly, _ := rectProps(10.8, 5.6) // Perimeter is discarded
	fmt.Printf("Area (with Blank Identifier) %f\n", areaOnly)
}
