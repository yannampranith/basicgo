package main

import "fmt"

func main() {
	number := 20
	case1(number)
	fmt.Println("Address of number in main: ", &number)
	fmt.Println("After case1, In main:", number)
	case2(&number)
	fmt.Println("After case2, In main:", number)
}

//Pass by value

func case1(number int) {
	number = number + 2
	fmt.Println("In case1, value is: ", number)
	fmt.Println("In case2, address is: ", number)
}

// Pass by reference
func case2(number *int) {
	*number = *number + 2
	fmt.Println("In case2, value is:", number)
	fmt.Println("In case2, address is :", &number)
}