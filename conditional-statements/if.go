package main

import "fmt"

func main() {
	name := "Manish"
	age := 12

	if age >= 18 {
		fmt.Println(name, "is and adult!")
	} else {
		fmt.Println(name, "is not an adult!")
		fmt.Println(name, "is learning to code!")
	}
}