package main

import (
	"fmt"
)

type Car struct {
	wheels int
	color  string
	cost   int
}

//Embedded structures
type Vehicle struct {
	Car
}

func main() {

	toyota := Car{
		wheels: 4,
		color:  "white",
		cost:   400000,
	}
	fmt.Println(toyota)

	//Anonymous  Structures

	tesla := struct {
		Manufacturer string
		Model        string
	}{
		Manufacturer: "Tesla",
		Model:        "V3",
	}

	fmt.Println(tesla.Manufacturer)

	v := Vehicle{}
	v.wheels = 4
	fmt.Println(v)
}
