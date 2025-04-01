package main

import (
	"fmt"
	"time"
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

    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the Weekend!")
    default:
        fmt.Println("It's a Weekday!")
    }

    user := "Manish"
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("Good Morning!", user)
    default: 
        fmt.Println("Good Afternoon!", user)
    }

    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a boolean")
        case int: 
            fmt.Println("I'm an integer!")
        default:
            fmt.Println("Don't know type \n", t)
        }
    }
    
    whatAmI(true)
    whatAmI(2)
    whatAmI("Hello, World!")
}