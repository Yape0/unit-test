package logic_02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal_04(t *testing.T) {
	result := SoalNo4(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[9][9], 253)
}
