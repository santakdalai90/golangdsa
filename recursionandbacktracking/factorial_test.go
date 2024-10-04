package recursionandbacktracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	assert.Equal(t, 120, Factorial(5))
	assert.Equal(t, 720, Factorial(6))
	assert.Equal(t, 120, FactorialIterative(5))
	assert.Equal(t, 720, FactorialIterative(6))

}
