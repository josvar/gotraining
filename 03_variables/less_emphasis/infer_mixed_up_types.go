package main

import "fmt"

func main() {
	var message = "Hello world"
	var a, b, c = 1, false, '\x61'

	fmt.Println(message, a, b, string(c))
}
