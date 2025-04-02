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

	secArr := [2][3]int{
		{1,2,3},
		{1,2,3},
	}
	fmt.Println(secArr)
}