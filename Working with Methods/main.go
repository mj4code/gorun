package main

import (
	"fmt"
)

func getPoint() (int, int) {
	return 2, 3
}

//Give name to return type variable
func getAnotherPoint() (x, y int) {
	return
}
func main() {
	x, y := getPoint()
	//x, y := getAnotherPoint()

	fmt.Printf("The x coordinate is %v and y coordinate is %v", x, y)

}
