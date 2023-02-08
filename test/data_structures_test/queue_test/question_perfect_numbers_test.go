package queue_test

import (
	"golang-dsa/src/data_structures/queue"
	"testing"

	"github.com/stretchr/testify/require"
)

// go test -v ./test/data_structures_test/queue_test/...

func Test_GetPerfectNumber(t *testing.T) {
	n := 4
	require.Equal(t, "1221", queue.GetPerfectNumber(n))

	n = 21
	require.Equal(t, "12211221", queue.GetPerfectNumber(n))
}
