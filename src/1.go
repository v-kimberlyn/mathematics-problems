package main

import "math/rand"

func main() {
	fmt.Println(generateMathProblem())
}

func generateMathProblem() string {
	// Generate a random number between 1 and 10
	num := rand.Intn(10) + 1

	// Generate a random math operation (+, -, *, /)
	op := "+"
	switch rand.Intn(4) {
	case 0:
		op = "+"
	case 1:
		op = "-"
	case 2:
		op = "*"
	case 3:
		op = "/"
	}

	// Generate a random answer
	answer := num * 5

	// Generate the math problem as a string
	return fmt.Sprintf("%d %s %d", num, op, answer)
}
