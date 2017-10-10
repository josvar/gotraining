package main

import (
	"fmt"
)

func main() {
	switch "Josh" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Josh":
		fmt.Println("Wassup Josh")
	case "Less":
		fmt.Println("Wassup Less")
	default:
		fmt.Println("Have you no friends?")
	}
}

/*
  no default fallthrough
  fallthrough is optional
  -- you can specify fallthrough by explicitly stating it
  -- break isn't needed like in other languages
*/
