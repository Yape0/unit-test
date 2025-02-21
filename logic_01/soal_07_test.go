package logic_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal_07(t *testing.T) {
	result := Soal_07(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[4], 9)
	assert.NotEqual(t, result[3], 9)

	result = Soal_07(11)
	assert.Equal(t, len(result), 11)
	assert.Equal(t, result[5], 11)
	assert.NotEqual(t, result[3], 9)

}
