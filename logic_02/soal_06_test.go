package logic_02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoalNo6(t *testing.T) {
	result := SoalNo6(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[8][8], 205)

}
