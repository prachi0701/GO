package main

import (
    "fmt"
)

// Variadic function that finds a number in a list of numbers
func find(num int, nums ...int) {
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
            break
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}

// Function that changes the first element of a string slice
func change(s ...string) {
    s[0] = "Go"
    s = append(s, "playground")
    fmt.Println("Inside change:", s)
}

func main() {
    // Using variadic function find
    find(89, 89, 90, 95)
    find(45, 56, 67, 45, 90, 109)
    find(78, 38, 56, 98)
    find(87) // Passing no variadic arguments

    // Passing a slice to a variadic function
    nums := []int{89, 90, 95}
    find(89, nums...) // Using ... to pass a slice to variadic function

    // Demonstrating the effect of modifying a slice inside a variadic function
    welcome := []string{"hello", "world"}
    change(welcome...)
    fmt.Println("Outside change:", welcome) // Slice is modified
}

