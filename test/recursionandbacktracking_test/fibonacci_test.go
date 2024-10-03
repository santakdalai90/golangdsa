package recursionandbacktracking_test

import (
	"testing"

	"github.com/santakdalai90/golangdsa/src/recursionandbacktracking"
	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	assert.Equal(t, 8, recursionandbacktracking.Fibonacci(6))
	assert.Equal(t, 8, recursionandbacktracking.FibonacciIterative(6))

	for i := 0; i < 100; i++ {
		assert.Equal(t, recursionandbacktracking.Fibonacci(i), recursionandbacktracking.FibonacciIterative(i))
	}
}
