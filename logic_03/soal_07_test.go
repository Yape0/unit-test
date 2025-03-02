package logic_03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoalNo07(t *testing.T) {
	result := SoalNo07(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[0][4], 1)
	assert.Equal(t, result[4][8], 33)
	assert.Equal(t, result[4][4], 41)
	assert.Equal(t, result[4][0], 49)
}
