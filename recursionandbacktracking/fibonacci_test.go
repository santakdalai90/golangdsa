package recursionandbacktracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	assert.Equal(t, 8, Fibonacci(6))
	assert.Equal(t, 8, FibonacciIterative(6))

	for i := 0; i < 100; i++ {
		assert.Equal(t, Fibonacci(i), FibonacciIterative(i))
	}
}
