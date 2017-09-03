package main

import "fmt"

const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

const d = iota // 0
const e = iota // 0

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d) // 0
	fmt.Println(e) // 0
}
