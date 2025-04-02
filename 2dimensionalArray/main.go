package main

import "fmt"

func main() {

	names := [...]string{"a", "b", "c", "d", "e", "f"}

	for i := 1; i < len(names); i++ {
		for j := 1; j < i; j++ {
			fmt.Print(names[i])
		}
		fmt.Println()
	}

}