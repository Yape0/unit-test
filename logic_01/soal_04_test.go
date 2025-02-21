package logic_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoalNo04(t *testing.T) {

	result := Soal_04(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[3], 13)
	assert.NotEqual(t, result[3], 12)

	result = Soal_04(13)
	assert.NotEqual(t, result[3], 9)

}
