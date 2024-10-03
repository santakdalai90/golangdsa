package recursionandbacktracking_test

import (
	"testing"

	"github.com/santakdalai90/golangdsa/src/recursionandbacktracking"
	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	assert.Equal(t, 120, recursionandbacktracking.Factorial(5))
	assert.Equal(t, 720, recursionandbacktracking.Factorial(6))
	assert.Equal(t, 120, recursionandbacktracking.FactorialIterative(5))
	assert.Equal(t, 720, recursionandbacktracking.FactorialIterative(6))

}
