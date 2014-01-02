package main

import (
	"fmt"
)

func add(x, y, z int, s string, a int) int {
	fmt.Println(a, s)
	return x + y
}

func main() {
	fmt.Println(add(42, 13, 55, "foo", 123))
}
