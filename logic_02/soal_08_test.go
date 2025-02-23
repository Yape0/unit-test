package logic_02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoalNo8(t *testing.T) {
	result := SoalNo8(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[8][8], 0)

}
