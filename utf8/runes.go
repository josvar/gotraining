package main

import "fmt"

func main() {
	c := '⌘'
	d := `\n \t`
	fmt.Printf("%v \n", c)
	fmt.Printf("%#v \n", c)
	fmt.Printf("%#X \n", c)
	fmt.Printf("%q \n", c)
	fmt.Printf("%c \n", c)

	fmt.Printf("%v", d)
}
