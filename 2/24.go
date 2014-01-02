package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 100; i++ {
		if zz := z - (z*z-x)/(2*z); math.Abs(zz-z) < 1e-10 {
			return z
		} else {
			z = zz
			fmt.Println(i)
		}

	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
