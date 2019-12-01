package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuel12(t *testing.T) {
	mass := 12
	fuel := calculateFuel(mass)
	assert.True(t, fuel == 2)
}

func TestFuel14(t *testing.T) {
	mass := 14
	fuel := calculateFuel(mass)
	assert.True(t, fuel == 2)
}

func TestFuel1969(t *testing.T) {
	mass := 1969
	fuel := calculateFuel(mass)
	assert.True(t, fuel == 654)
}

func TestFuel100756(t *testing.T) {
	mass := 100756
	fuel := calculateFuel(mass)
	assert.True(t, fuel == 33583)
}
