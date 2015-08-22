package svector

import "strconv"

// Vector is a representation of a sparse vector in some
// n-dimensional Euclidean space.
type Vector struct {
	data map[string]float64
}

// Get the vector's component associated to a particular key string.
func (v *Vector) Get(key string) (float64, bool) {
	value, prs := v.data[key]
	return value, prs
}

// Set a vector's component.  Zero values are ignored.
func (v *Vector) Set(key string, value float64) {
	if value > 0 {
		v.data[key] = value
	}
}

// Norm of the vector (with respect to the standard Euclidean inner product).
func (v *Vector) Norm() float64 {
	return Norm(v)
}

// New vector ready to have its components specified.
func New() *Vector {
	return &Vector{data: make(map[string]float64)}
}

// Make a vector from a map representation.
func Make(input interface{}) *Vector {
	var v *Vector

	switch t := input.(type) {
	case map[string]float64:
		v = makeFromMap(t)
	case []float64:
		v = makeFromSlice(t)
	}

	return v
}

func makeFromMap(input map[string]float64) *Vector {
	v := New()

	for k, vk := range input {
		v.Set(k, vk)
	}

	return v
}

func makeFromSlice(input []float64) *Vector {
	v := New()

	for k, vk := range input {
		v.Set(strconv.Itoa(k), vk)
	}

	return v
}
