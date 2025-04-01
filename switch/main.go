package main

import (
	"fmt"
)
func main() {
    i := 12
    switch i {
    case 1: 
    fmt.Println("One")
    case 2: 
        fmt.Println("Two")
    case 3: 
        fmt.Println("Three")
    default: 
        fmt.Println(i, "Not found!")
    }
}