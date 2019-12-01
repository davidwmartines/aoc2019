package main

import "fmt"
import "math"

func main() {
	fmt.Println("hello!")
}

func calculateFuel(mass int) int {
	return int(math.Round(float64(mass/3)) - 2)
}
