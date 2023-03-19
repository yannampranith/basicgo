package main

import (
	"fmt"
	"math/big"
)

var (
	// Create a huge number by using the math/big package.
	Big = new(big.Int).Exp(big.NewInt(2), big.NewInt(100), nil)
	// Divide it by 10^99 to obtain a small integer value.
	Small = new(big.Int).Rsh(Big, 99)
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(int(Small.Int64())))
	fmt.Println(needFloat(float64(Small.Int64())))
	fmt.Println(needFloat(Big.Float64()))
}
This implementation uses the math/big package to perform the large integer calculations. The Big constant is defined as a new big.Int value that is created by raising 2 to the power of 100 using the Exp method. The Small constant is defined as the result of right-shifting the Big value by 99 bits using the Rsh method. The needInt and needFloat functions are the same as before, but they now require explicit type conversions from big.Int to int64 or float64.





