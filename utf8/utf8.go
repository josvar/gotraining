package main

import "fmt"

func main() {
	c := '\U00010348'
	d := '€'
	e := '\U0010FFFF'
	fmt.Printf("%c \n", c)
	fmt.Printf("%x \n", d)
	fmt.Printf("%c \n", e)
}
