package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	max := 10
	min := 0
	result := rand.Intn(max-min+1) + min
	fmt.Println("Result:", result)
}
