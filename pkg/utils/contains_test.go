package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMapKey(t *testing.T) {
	containsInstance := ContainsMapKey("hello", map[string]string{
		"hello": "ah",
		"world": "ahh",
	})
	assert.True(t, containsInstance)

	notFoundInstance := ContainsMapKey("hello", map[string]string{
		"ah":  "hello",
		"ahh": "world",
	})
	assert.False(t, notFoundInstance)
}
