package logic_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoalNo02(t *testing.T) {

	result := Soal_02(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[3], 8)
	assert.NotEqual(t, result[3], 9)

	result = Soal_02(13)
	assert.NotEqual(t, result[3], 9)

}
