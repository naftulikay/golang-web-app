package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNullBytes32(t *testing.T) {
	var comparison [32]byte

	for index, _ := range comparison {
		// force it
		comparison[index] = 0
	}

	assert.Equal(t, comparison, NullBytes32())
}

func TestNullBytes64(t *testing.T) {
	var comparison [64]byte

	for index, _ := range comparison {
		// force it
		comparison[index] = 0
	}

	assert.Equal(t, comparison, NullBytes64())
}
