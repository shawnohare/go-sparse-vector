package svector

import "math"

// Add two sparse vectors to obtain a new sparse vector. This does not
// mutate the two input vectors.
func Add(v, w *Vector) *Vector {
	x := New()
	for k, vk := range v.data {
		wk, _ := w.Get(k)
		x.Set(k, vk+wk)
	}

	return x
}

// Scale a vector by a given scalar.  This does not mutate the input vector.
func Scale(scalar float64, v *Vector) *Vector {
	x := New()
	for k, vk := range v.data {
		x.Set(k, scalar*vk)
	}

	return x
}

// Dot product (standard Euclidean inner product) between two vectors.
func Dot(v, w *Vector) float64 {
	sum := 0.0

	for k, vk := range v.data {
		wk, _ := w.Get(k)
		sum += vk * wk
	}

	return sum
}

// Norm of a vector with respect to the standard Euclidean inner product.
func Norm(v *Vector) float64 {
	return math.Sqrt(Dot(v, v))
}
