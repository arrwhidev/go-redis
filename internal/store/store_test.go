package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func NewEntryWithoutExpiry(value string) *Entry {
	return NewEntry(value, -1)
}

func TestSingleton(t *testing.T) {
	store := Instance()
	assert.Equal(t, store, Instance())
}

func TestSetReturnsError_whenUnknownKey(t *testing.T) {
	store := NewStore()
	_, err := store.Get("hello")
	assert.NotNil(t, err)
}

func TestSetOverwritesValue_whenKeyExists(t *testing.T) {
	store := NewStore()
	store.Set("hello", NewEntryWithoutExpiry("world"))
	store.Set("hello", NewEntryWithoutExpiry("world2"))
	e, _ := store.Get("hello")
	assert.Equal(t, "world2", e.Value)
}

func TestItCanSetAndGet(t *testing.T) {
	store := NewStore()
	store.Set("hello", NewEntry("world", -1))
	e, _ := store.Get("hello")
	assert.Equal(t, "world", e.Value)
}

func TestGetReturnsEntry_whenNotExpired(t *testing.T) {
	store := NewStore()
	future := time.Now().Add(5 * time.Second).UnixNano()
	store.Set("hello", NewEntry("world", future))
	e, _ := store.Get("hello")
	assert.Equal(t, future, e.Expires)
}

func TestGetReturnsEntry_whenExpiryIsMinus1(t *testing.T) {
	store := NewStore()
	entry := NewEntry("world", -1)
	store.Set("hello", entry)
	e, _ := store.Get("hello")
	assert.Equal(t, entry, e)
}

func TestGetReturnsNil_whenExpired(t *testing.T) {
	store := NewStore()
	past := time.Now().Add(-5 * time.Second).UnixNano()
	store.Set("hello", NewEntry("world", past))
	e, _ := store.Get("hello")
	assert.Nil(t, e)
}

func TestKeysReturnsAllNonExpiredKeys(t *testing.T) {
	store := NewStore()
	store.Set("hello1", NewEntry("world", -1))
	store.Set("hello2", NewEntry("world", -1))
	store.Set("hello3", NewEntry("world", -1))
	store.Set("expired", &Entry{"world", time.Now().Add(-5 * time.Second).UnixNano()})

	keys := store.Keys()
	assert.Len(t, keys, 3)
	assert.Contains(t, keys, "hello1", "hello2", "hello3")
}
