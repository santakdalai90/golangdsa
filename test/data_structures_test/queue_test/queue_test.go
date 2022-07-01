package queue_test

import (
	"golang-dsa/src/data_structures/queue"
	"testing"
    "github.com/stretchr/testify/require"
)

// go test -v ./test/data_structures_test/queue_test/...

func Test_Queue(t *testing.T) {
	q := queue.New()
    q.Push("apple")
    q.Push(24837)
    q.Push(23.5645)
    q.Push("tiger")
    q.Print()

    v, err := q.Pop()
    require.NoError(t, err)
    t.Log(v)
    
}