package svector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	m := make(map[string]float64)
	v := New()
	assert.Equal(t, m, v.data)
}

func TestMakeFromMap(t *testing.T) {
	m := map[string]float64{"a": 1.0, "b": 2.0}
	v := Make(m)
	assert.Equal(t, m, v.data)
}

func TestMakeFromSlice(t *testing.T) {
	m := map[string]float64{"0": 1.0, "1": 2.0}
	s := []float64{1.0, 2.0}
	v := Make(s)
	assert.Equal(t, m, v.data)
}

func TestMakeFromJunk(t *testing.T) {
	v := Make(1.0)
	assert.Nil(t, v)
}

func TestSet(t *testing.T) {
	v := New()
	// Test that 0-values are not set.
	v.Set("key", 0)
	assert.Len(t, v.data, 0)

	v.Set("key", 1)
	assert.Len(t, v.data, 1)
	assert.Equal(t, 1.0, v.data["key"])
}

func TestGet(t *testing.T) {
	v := New()
	_, prs := v.Get("key")
	assert.False(t, prs)

	v.Set("key", 1)
	val, prs := v.Get("key")
	assert.True(t, prs)
	assert.Equal(t, 1.0, val)
}
