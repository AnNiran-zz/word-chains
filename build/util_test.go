package build

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// 
func TestOrderInput(t *testing.T) {
	inputProperOrder := []string{
		"alphabet",
		"connection",
	}

	inputReverseOrder := []string{
		"zoo",
		"brighter",
	}

	resProper := orderInput(inputProperOrder[0], inputProperOrder[1])
	expectedPropOrder := true

	resReverse := orderInput(inputReverseOrder[0], inputReverseOrder[1])
	expectedReverseOrder := false

	assert.Equal(t, resProper, expectedPropOrder)
	assert.Equal(t, resReverse, expectedReverseOrder)
}
