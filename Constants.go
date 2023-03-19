package main

import "fmt"

const Greeting = "Bonjour"

func main() {
	const Name = "Alice"
	fmt.Println(Greeting, Name)
	fmt.Println("The value of pi is approximately", Pi)

	const Answer = 42
	fmt.Println("The answer to everything is", Answer)
}
