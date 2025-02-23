package logic_02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoalNo5(t *testing.T) {
	result := SoalNo5(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[8][8], 161)
}
