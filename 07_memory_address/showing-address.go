package main

import "fmt"

func main() {
	a := 43
	b := &a

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
	fmt.Printf("DECIMAL\t%d\t%d\n", &a, b)
	fmt.Printf("HEX\t\t%#X\t%#X\n", &a, b)
}
