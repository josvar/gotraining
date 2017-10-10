package main

import (
	"fmt"
)

func main() {
	switch "Josh" {
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Josh":
		fmt.Println("Wassup Josh")
		fallthrough
	case "Medhi":
		fmt.Println("Wassup Medhi")
		fallthrough
	case "Julian":
		fmt.Println("Wassup Julian")
	case "Less":
		fmt.Println("Wassup Less")
	}
}
