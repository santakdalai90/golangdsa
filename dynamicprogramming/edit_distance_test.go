package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEditDistance(t *testing.T) {
	tests := []struct {
		name         string
		source       string
		target       string
		editDistance int
	}{
		{"empty strings", "", "", 0},
		{"identical strings", "kitten", "kitten", 0},
		{"replacements", "kitten", "sitting", 3},
		{"deletion", "sitting", "sittng", 1},
		{"insertion and replacement", "saturday", "sunday", 3},
		{"long string", "intention", "execution", 5},
		{"case insensitive", "sunday", "Sunday", 1},
		{"unicode characters", "café", "cafe", 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			editDistance := EditDistance(test.source, test.target)
			assert.Equal(t, test.editDistance, editDistance)
		})
	}
}
