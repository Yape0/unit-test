package logic_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal_09(t *testing.T) {
	result := Soal_09(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[3], 12)
	assert.NotEqual(t, result[3], 11)

	result = Soal_09(11)
	assert.Equal(t, len(result), 11)
	assert.Equal(t, result[5], 18)
	assert.Equal(t, result[10], 3)

}
