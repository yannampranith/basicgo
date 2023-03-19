package main

import "fmt"

func swap(x *string, y *string) {
	*x, *y = *y, *x
}

func main() {
	a := "hello"
	b := "world"
	swap(&a, &b)
	fmt.Println(a, b)
}
