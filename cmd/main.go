package main

import (
	"fmt"

	mymath "github.com/tjalfe/godemo/pkg"
)

func main() {
	var x int64 = 10
	var y int64 = 4

	fmt.Printf("%v + %v = %v\n", x, y, mymath.Add(x, y))
	fmt.Printf("%v - %v = %v\n", x, y, mymath.Subtract(x, y))
	fmt.Printf("%v / %v = %.2f\n", x, y, mymath.Divide(x, y))

}
