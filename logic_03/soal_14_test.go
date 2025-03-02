package logic_03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoalNo14(t *testing.T) {
	result := SoalNo14(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[0][4], 137)
	assert.Equal(t, result[4][8], 281)
	assert.Equal(t, result[4][4], 145)
	assert.Equal(t, result[4][0], 9)
}
