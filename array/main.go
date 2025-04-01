package main

import "fmt"

func main() {
	a := [...]string{"apple", "banana", "mango"}
	fmt.Println(len(a))
	fmt.Println(a)

	num := [...]int{2,24: 34, 27: 5}
	fmt.Println(num)
}