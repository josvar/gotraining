package main

import (
	"fmt"
	"gotraining/02_packages/utils"
	"gotraining/02_packages/icomefrombsas"
)

func main() {
	fmt.Println(utils.MyVar)

	//cannot be accessed
	//fmt.Println(utils.anotherVar)

	fmt.Println(icomefrombsas.BearName)
	fmt.Println(utils.Reverse("!oG ,olleH"))
}
