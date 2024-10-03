package queue_test

import (
	"testing"

	"github.com/santakdalai90/golangdsa/src/datastructures/queue"

	"github.com/stretchr/testify/require"
)

// go test -v ./test/data_structures_test/queue_test/...

func Test_GetPerfectNumber(t *testing.T) {
	n := 4
	require.Equal(t, "1221", queue.GetPerfectNumber(n))

	n = 21
	require.Equal(t, "12211221", queue.GetPerfectNumber(n))
}
