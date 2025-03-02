package logic_03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoalNo02(t *testing.T) {
	result := SoalNo02(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[8][8], 89)
	assert.Equal(t, result[7][8], 85)
}
