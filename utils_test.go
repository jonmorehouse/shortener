package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
  assert.Equal(t, factorial(0), 1)
  assert.Equal(t, factorial(1), 1)
  assert.Equal(t, factorial(2), 2)
  assert.Equal(t, factorial(3), 6)
}

func TestPossiblePermutations(t *testing.T) {
  assert.Equal(t, possiblePermutations(3, 6), 120)
}

func TestIterator(t *testing.T) {
  iterate()
}

func TestBuildString(t *testing.T) {

  buildString()
}


