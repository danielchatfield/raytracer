package raytracer

import "math"

// Vector is a struct representing a floating point vector.
type Vector struct {
	X, Y, Z float64
}

type Ray struct {
	Position  Vector
	Direction Vector
}

// Add computes the sum of two vectors
func (v *Vector) Add(v2 *Vector) *Vector {
	return &Vector{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
		Z: v.Z + v2.Z,
	}
}

// Subtract takes the passed vector away from the current one and returns a pointer
func (v *Vector) Subtract(v2 *Vector) *Vector {
	return &Vector{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
		Z: v.Z - v2.Z,
	}
}

// Scale scales the vector by the specified amount
func (v *Vector) Scale(s float64) *Vector {
	return &Vector{
		X: v.X * s,
		Y: v.Y * s,
		Z: v.Z * s,
	}
}

// DotProduct computes the vector dot product of the two vectors
func (v *Vector) DotProduct(v2 *Vector) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v *Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vector) UnitVector() *Vector {
	return v.Scale(1.0 / v.Length())
}
