package logic_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal_06(t *testing.T) {
	result := Soal_06(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[3], 21)
	assert.NotEqual(t, result[3], 24)

	result = Soal_06(13)
	assert.NotEqual(t, len(result), 12)
}
