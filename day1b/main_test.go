package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuel12(t *testing.T) {
	mass := 12
	fuel := calculateFuel(mass)
	assert.Equal(t, fuel, 2)
}

func TestFuel14(t *testing.T) {
	mass := 14
	fuel := calculateFuel(mass)
	assert.Equal(t, fuel, 2)
}

func TestFuel1969(t *testing.T) {
	mass := 1969
	fuel := calculateFuel(mass)
	assert.Equal(t, fuel, 654)
}

func TestFuel100756(t *testing.T) {
	mass := 100756
	fuel := calculateFuel(mass)
	assert.Equal(t, fuel, 33583)
}

func TestFuelAdjusted14(t *testing.T) {
	mass := 14
	fuel := calculateFuelAdjusted(mass)
	assert.Equal(t, fuel, 2)
}

func TestFuelAdjusted1969(t *testing.T) {
	mass := 1969
	fuel := calculateFuelAdjusted(mass)
	assert.Equal(t, fuel, 966)
}

func TestFuelAdjusted100756(t *testing.T) {
	mass := 100756
	fuel := calculateFuelAdjusted(mass)
	assert.Equal(t, fuel, 50346)
}
