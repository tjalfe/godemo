package main

import (
	"fmt"
	"log"

	mymath "github.com/tjalfe/godemo/pkg"
)

func main() {
	var x int64 = 10
	var y int64 = 0

	fmt.Printf("%v + %v = %v\n", x, y, mymath.Add(x, y))
	fmt.Printf("%v - %v = %v\n", x, y, mymath.Subtract(x, y))

	div, err := mymath.Divide(x, y)
	if err != nil {
		log.SetPrefix("mymath: ")
		log.Fatal(err)
	}

	fmt.Printf("%v / %v = %.2f\n", x, y, div)

}
