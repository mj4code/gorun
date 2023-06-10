package main

import (
	"fmt"
)

func main() {

	str := "Hello Go !"

	if length := len(str); length > 3 {
		fmt.Println("Length is greater than 3")
	} else {
		fmt.Println("Too short !")
	}
}
