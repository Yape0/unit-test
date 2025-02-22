package logic_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal_10(t *testing.T) {
	result := Soal_10(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[1], "fizz")
	assert.Equal(t, result[2], 4)
	assert.NotEqual(t, result[1], 3)

	result = Soal_10(12)

	assert.Equal(t, len(result), 12)
	assert.Equal(t, result[11], "fizz")
	assert.Equal(t, result[10], 64)
}
