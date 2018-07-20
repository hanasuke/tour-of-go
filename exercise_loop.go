package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for ; z*z < x; {
		z += 1
	}

	for p := 0; p < 10; p++ {
		 z = (z + x/z)/2
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
