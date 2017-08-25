package main

import "fmt"

func main() {
	x := 0
	var increment = func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(y)
}
var y = 99
