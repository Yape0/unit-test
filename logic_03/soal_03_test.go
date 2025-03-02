package logic_03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoalNo3(t *testing.T) {
	result := SoalNo03(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[0][0], 1)
	assert.Equal(t, result[8][0], 125)
	assert.Equal(t, result[7][0], 121)
}
