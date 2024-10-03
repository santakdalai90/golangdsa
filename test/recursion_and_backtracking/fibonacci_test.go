package recursion_and_backtracking

import (
	"testing"

	"github.com/santakdalai90/GolangDSA/src/recursion_and_backtracking"
	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	assert.Equal(t, 8, recursion_and_backtracking.Fibonacci(6))
	assert.Equal(t, 8, recursion_and_backtracking.FibonacciIterative(6))

	for i := 0; i < 100; i++ {
		assert.Equal(t, recursion_and_backtracking.Fibonacci(i), recursion_and_backtracking.FibonacciIterative(i))
	}
}
