package main

import "fmt"

func add(x *int, y *int) int {
	return *x + *y
}

func main() {
	x := 42
	y := 13
	fmt.Println(add(&x, &y))
}
