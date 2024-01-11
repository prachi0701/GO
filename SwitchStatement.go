package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Example of basic switch usage
	fingerExample()

	// Example of a switch with a default case
	defaultCaseExample()

	// Example of multiple expressions in a case
	multipleExpressionsExample()

	// Example of an expressionless switch
	expressionlessSwitchExample()

	// Example of switch with fallthrough
	fallthroughExample()

	// Example of breaking out of a switch
	breakExample()

	// Example of generating a random even number and breaking the outer loop
	randomEvenNumberExample()
}

func fingerExample() {
	finger := 4
	fmt.Printf("Finger %d is ", finger)
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	}
}

func defaultCaseExample() {
	switch finger := 8; finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("Incorrect finger number")
	}
}

func multipleExpressionsExample() {
	letter := "i"
	fmt.Printf("Letter %s is a ", letter)
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}
}

func expressionlessSwitchExample() {
	num := 75
	switch {
	case num >= 0 && num <= 50:
		fmt.Printf("%d is greater than 0 and less than 50\n", num)
	case num >= 51 && num <= 100:
		fmt.Printf("%d is greater than 51 and less than 100\n", num)
	}
}

func fallthroughExample() {
	num := 75
	switch {
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
	}
}

func breakExample() {
	num := -5
	switch {
	case num < 50:
		if num < 0 {
			break
		}
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
	}
}

func randomEvenNumberExample() {
randloop:
	for {
		switch i := rand.Intn(100); {
		case i%2 == 0:
			fmt.Printf("Generated even number %d\n", i)
			break randloop
		}
	}
}
