package main

import "fmt"

func mutate(s string) string {
    runeSlice := []rune(s)
    runeSlice[0] = 'a' // Replace the first rune
    return string(runeSlice)
}

func main() {  
    h := "hello"
    fmt.Println(mutate(h))
}
