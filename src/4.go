package main

import "math/rand"

func main() {
	rand.Seed(9)
	x := rand.Float64() * 10
	y := rand.Float64() * 10
	z := x + y
	fmt.Println(z)
}
