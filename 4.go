package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	now := time.Now().UnixNano()
	random := rand.New(rand.NewSource(now))
	fmt.Println("My favorite number is", random.Intn(10))
}
