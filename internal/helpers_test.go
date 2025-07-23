package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIn(t *testing.T) {

	found, pos := IsIn(1, []int{1, 2, 3})

	assert.True(t, found)
	assert.Equal(t, pos, 0)

	found, pos = IsIn(1, []int{4, 1, 3})

	assert.True(t, found)
	assert.Equal(t, pos, 1)

	found, pos = IsIn(3, []int{1, 2, 3, 3, 3})

	assert.True(t, found)
	assert.Equal(t, pos, 2)
}

func TestNotIsIn(t *testing.T) {

	found, pos := IsIn(4, []int{1, 2, 3})

	assert.False(t, found)
	assert.Equal(t, pos, -1)

	found, pos = IsIn(-5, []int{5, 2, 3})

	assert.False(t, found)
	assert.Equal(t, pos, -1)
}
