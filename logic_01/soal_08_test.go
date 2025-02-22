package logic_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal_08(t *testing.T) {
	result := Soal_08(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[9], 2)
	assert.NotEqual(t, result[4], 9)

	result = Soal_08(11)

	assert.Equal(t, result[6], 10)
	assert.Equal(t, result[4], 10)
	assert.Equal(t, result[8], 6)
	assert.NotEqual(t, len(result), 10)

}
