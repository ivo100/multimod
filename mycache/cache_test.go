package mycache

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCache_Get(t *testing.T) {
	cache := New(time.Second)
	cache.Set("key", "value", time.Second)
	value, exists := cache.Get("key")
	assert.True(t, exists)
	assert.Equal(t, "value", value)
}
