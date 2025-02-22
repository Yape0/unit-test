package logic_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal_12(t *testing.T) {
	result := soal_12(12)

	assert.Equal(t, len(result), 12)
	assert.Equal(t, result[3], 7)
	assert.NotEqual(t, result[8], 5)

	result = soal_12(14)

	assert.NotEqual(t, result[10], 8)

}
