package main

import (
	"fmt"
)

func main() {

	foo := "hello world"

	chars := []int32(foo)

	fmt.Printf("%c\n", chars)
	fmt.Printf("%v\n", chars)
	fmt.Printf("%b\n", chars)
	fmt.Printf("%x\n", chars)
}
