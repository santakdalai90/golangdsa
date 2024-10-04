package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckSpelling(t *testing.T) {
	sc, err := NewSpellChecker(
		"sherlock_homes.txt",
		1,
		2,
		5,
	)

	assert.NoError(t, err)

	tests := []struct {
		name        string
		word        string
		corrections []string
		err         error
	}{
		{"correct word", "information", []string{"information"}, nil},
		{"incorrect word 1", "stroes", []string{"stores", "strokes"}, nil},
		{"incorrect word 2", "primise", []string{"promise", "premise"}, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			corrections, err := sc.CheckSpelling(test.word)
			if test.err != nil {
				assert.ErrorIs(t, test.err, err)
			}
			assert.Equal(t, test.corrections, corrections)
		})
	}
}
