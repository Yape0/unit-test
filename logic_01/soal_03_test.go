package logic_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoalNo03(t *testing.T) {

	result := Soal_03(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[3], 12)
	assert.NotEqual(t, result[3], 13)

	result = Soal_03(13)
	assert.NotEqual(t, result[3], 9)

}
