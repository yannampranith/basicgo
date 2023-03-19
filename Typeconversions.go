package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := 3, 4
	f := math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	fmt.Println(x, y, z)
}
This code uses the short form variable declaration syntax name := value to initialize the variables x and y. It also uses type inference to automatically determine the type of the variables f



