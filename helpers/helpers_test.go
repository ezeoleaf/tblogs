package helpers

import (
	"testing"

	"github.com/ezeoleaf/tblogs/models"
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

func TestIsHash(t *testing.T) {

	values := []models.Post{
		models.Post{Title: "Post1", Hash: "H1"},
		models.Post{Title: "Post2", Hash: "H2"},
		models.Post{Title: "Post3", Hash: "H3"}}

	found, pos := IsHash("H1", values)

	assert.True(t, found)
	assert.Equal(t, pos, 0)

	found, pos = IsHash("H3", values)

	assert.True(t, found)
	assert.Equal(t, pos, 2)
}

func TestNotIsHash(t *testing.T) {

	values := []models.Post{
		models.Post{Title: "Post1", Hash: "H1"},
		models.Post{Title: "Post2", Hash: "H2"},
		models.Post{Title: "Post3", Hash: "H3"}}

	found, pos := IsHash("h1", values)

	assert.False(t, found)
	assert.Equal(t, pos, -1)

	found, pos = IsHash("H4", values)

	assert.False(t, found)
	assert.Equal(t, pos, -1)
}
