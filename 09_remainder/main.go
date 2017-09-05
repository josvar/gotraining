package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	y := rand.Intn(10)
	rem := y % 2

	fmt.Println(y)
	if rem == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}


}
