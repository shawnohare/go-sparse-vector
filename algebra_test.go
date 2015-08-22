package svector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	e := New()
	v := Make([]float64{1, 2})
	w := Add(v, v)

	expected := Make([]float64{2, 4})
	assert.Equal(t, expected.data, w.data)
	assert.Equal(t, w.data, Add(w, e).data)
}

func TestScale(t *testing.T) {
	e := New()
	v := Make([]float64{1, 2})
	w := Scale(3, v)
	expected := Make([]float64{3, 6})
	assert.Equal(t, expected.data, w.data)
	assert.Equal(t, e.data, Scale(0, w).data)
}

func TestDot(t *testing.T) {
	v := Make([]float64{1, 2})
	e := New()
	assert.Equal(t, 0.0, Dot(v, e))
	assert.Equal(t, 5.0, Dot(v, v))
}

func TestNorm(t *testing.T) {
	v := Make([]float64{3, 4})
	e := New()
	assert.Equal(t, 5.0, Norm(v))
	assert.Equal(t, 5.0, v.Norm())
	assert.Equal(t, 0.0, Norm(e))
	assert.Equal(t, 0.0, e.Norm())
}
