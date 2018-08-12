package lexorank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessEmptyPrevEmptyNext(t *testing.T) {
	rank, ok := Rank("", "")
	assert.Equal(t, "U", rank)
	assert.Equal(t, true, ok)
}

func TestSuccessEmptyPrev(t *testing.T) {
	rank, ok := Rank("", "2")
	assert.Equal(t, "1", rank)
	assert.Equal(t, true, ok)
}

func TestSuccessEmptyNext(t *testing.T) {
	rank, ok := Rank("x", "")
	assert.Equal(t, "y", rank)
	assert.Equal(t, true, ok)
}

func TestSuccessNewDigit(t *testing.T) {
	rank, ok := Rank("aaaa", "aaab")
	assert.Equal(t, "aaaaU", rank)
	assert.Equal(t, true, ok)
}

func TestSuccessMidValue(t *testing.T) {
	rank, ok := Rank("aaaa", "aaac")
	assert.Equal(t, "aaab", rank)
	assert.Equal(t, true, ok)
}

func TestSuccessNewDigitMidValue(t *testing.T) {
	rank, ok := Rank("az", "b")
	assert.Equal(t, "azU", rank)
	assert.Equal(t, true, ok)
}

func TestFailSamePrevNext(t *testing.T) {
	rank, ok := Rank("aaaa", "aaaa")
	assert.Equal(t, "aaaa", rank)
	assert.Equal(t, false, ok)
}

func TestFailAdjacent(t *testing.T) {
	rank, ok := Rank("a", "a0")
	assert.Equal(t, "a", rank)
	assert.Equal(t, false, ok)
}
