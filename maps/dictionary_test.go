package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is test"}

	got, _ := dictionary.Search("test")
	want := "this is test"

	assert.Equal(t, want, got)
}

func TestSearchFails(t *testing.T) {
	dictionary := Dictionary{"test": "this is test"}

	got, err := dictionary.Search("unknown")
	want := ErrKeyNotFound

	assert.ErrorIs(t, err, want)
	assert.Equal(t, "", got)
}

func TestAddThenSearch(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is test")

	got, err := dictionary.Search("test")
	want := "this is test"
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestAddExistingKey(t *testing.T) {
	dictionary := Dictionary{"test": "this is test"}
	err := dictionary.Add("test", "this is test")
	assert.ErrorIs(t, err, ErrKeyAlreadyExists)
}

func TestUpdateExisting(t *testing.T) {
	dictionary := Dictionary{"test": "this is test"}
	err := dictionary.Update("test", "updated")
	assert.NoError(t, err)

	got, err := dictionary.Search("test")
	assert.NoError(t, err)
	assert.Equal(t, "updated", got)
}

func TestUpdateNonExisting(t *testing.T) {
	dictionary := Dictionary{}
	err := dictionary.Update("test", "updated")
	assert.ErrorIs(t, err, ErrKeyNotFound)

}

func TestDeleteExisting(t *testing.T) {
	dictionary := Dictionary{"test": "this is test"}
	err := dictionary.Delete("test")
	assert.NoError(t, err)

	_, err = dictionary.Search("test")
	assert.ErrorIs(t, err, ErrKeyNotFound)
}

func TestDeleteNonExisting(t *testing.T) {
	dictionary := Dictionary{}
	err := dictionary.Delete("test")
	assert.ErrorIs(t, err, ErrKeyNotFound)
}
