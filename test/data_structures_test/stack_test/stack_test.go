package stack_test

import (
	"golang-dsa/src/data_structures/stack"
	"testing"
    "github.com/stretchr/testify/require"
)

// go test -v ./test/data_structures_test/stack_test/...

func Test_Stack(t *testing.T) {
	s := stack.New()
    s.Push("apple")
    s.Push(24837)
    s.Push(23.5645)
    s.Push("tiger")
    s.Print()

    v, err := s.Pop()
    require.NoError(t, err)
    t.Log(v)
    
}