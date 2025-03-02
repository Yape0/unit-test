package logic_03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoalNo05(t *testing.T) {
	result := SoalNo05(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[0][0], 1)
	assert.Equal(t, result[0][8], 1)
	assert.Equal(t, result[8][0], 1)
}
