package logic_03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoalNo13(t *testing.T) {
	result := SoalNo13(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[0][4], 9)
	assert.Equal(t, result[4][8], 1)
	assert.Equal(t, result[4][4], 9)
	assert.Equal(t, result[4][0], 1)
}
