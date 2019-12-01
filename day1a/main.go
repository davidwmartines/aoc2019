package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./day1a/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		fuel := calculateFuel(mass)
		total += fuel
	}
	fmt.Printf("total: %v\n", total)
}

func calculateFuel(mass int) int {
	return int(math.Round(float64(mass/3)) - 2)
}
