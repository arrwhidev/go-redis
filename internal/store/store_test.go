package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleton(t *testing.T) {
	store := Instance()
	assert.Equal(t, store, Instance())
}

func TestSetReturnsError_WhenUnknownKey(t *testing.T) {
	store := NewStore()
	_, err := store.Get("hello")
	assert.NotNil(t, err)
}

func TestSetOverwritesValue_WhenKeyExists(t *testing.T) {
	store := NewStore()
	store.Set("hello", "world")
	store.Set("hello", "world2")
	v, _ := store.Get("hello")
	assert.Equal(t, "world2", v)
}

func TestItCanSetAndGet(t *testing.T) {
	store := NewStore()
	store.Set("hello", "world")
	v, _ := store.Get("hello")
	assert.Equal(t, "world", v)
}
