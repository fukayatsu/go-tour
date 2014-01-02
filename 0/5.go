package main

import (
	"fmt"
	"math"
)

func main() {
	// http://golang.org/pkg/math/#Nextafter
	fmt.Printf("Now you have %g problems.",
		math.Nextafter(2, 3))
}
