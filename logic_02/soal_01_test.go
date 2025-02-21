package logic_02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoalNo01(t *testing.T) {
	result := SoalNo1(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[9][9], 19)
}
