package logic_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoalNo01(t *testing.T) {
	result := soalNo01(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[3], 7)
	assert.NotEqual(t, result[3], 9)

	result = soalNo01(13)
	assert.NotEqual(t, result[3], 9)
}
