package logic_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal_11(t *testing.T) {
	result := Soal_11(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[2], "buzz")
	assert.NotEqual(t, result[3], "fizz")

	result = Soal_11(11)
	assert.Equal(t, result[10], "buzz")
}
