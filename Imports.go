package main

import "fmt"

func main() {
	num := 7.0
	root := 1.0
	for i := 0; i < 10; i++ {
		root = (root + num/root) / 2
	}
	fmt.Printf("Now you have %g problems.\n", root)
}






