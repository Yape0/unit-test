package logic_02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoalNo9(t *testing.T) {
	result := SoalNo9(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[8][8], 17)
	assert.Equal(t, result[8][0], 1)
}
