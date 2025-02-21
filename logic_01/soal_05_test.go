package logic_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal_05(t *testing.T) {
	result := Soal_05(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[3], 14)
	assert.NotEqual(t, result[3], 15)

	result = Soal_05(13)
	assert.NotEqual(t, result[3], 15)
}
