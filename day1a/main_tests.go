package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuel(t *testing.T){
	mass := 12
	fuel = calculateFuel(mass)

	assert.True(t, fuel == 2)

}