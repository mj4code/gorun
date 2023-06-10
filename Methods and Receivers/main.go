package main

import (
	"fmt"
)

func (r rectangle) area() int {
	return r.breadth * r.length
}

type rectangle struct {
	length  int
	breadth int
}

func main() {

	r := rectangle{
		length:  4,
		breadth: 5,
	}

	fmt.Println("Area of the rectangle is : ", r.area())

}
