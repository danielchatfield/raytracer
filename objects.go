package raytracer

import (
	"image/color"
	"math"
)

// ObjectInterface is the interface for objects in the scene
type ObjectInterface interface {
	getCollision(*Vector) (bool, float64)
	getInterception(*Ray) (bool, *Vector)
	getColor(v *Vector) color.NRGBA
}

// Object represents an object in the scene
type Object struct {
	ObjectInterface
	Ambient  float64
	Diffuse  float64
	Specular float64
}

func (o Object) getColor(v *Vector) color.NRGBA {
	return color.NRGBA{255, 255, 255, 255}
}

// Light represents a light source
type Light struct {
}

// Sphere represents a spherical object
type Sphere struct {
	Object
	Centre Vector
	Radius float64
}

func (s Sphere) getInterception(r *Ray) (bool, *Vector) {
	a := (&r.Direction).DotProduct(&r.Direction)
	tmp := r.Position.Subtract(&s.Centre)
	b := (&r.Direction).Scale(2).DotProduct(tmp)
	c := tmp.DotProduct(tmp) - (s.Radius * s.Radius)

	d := b*b - 4*a*c

	if d < 0 {
		return false, &Vector{}
	}

	d = math.Sqrt(d)

	scale := (d - b) / (2 * a)

	return true, (&r.Direction).Scale(scale)
}

func (s Sphere) getCollision(v *Vector) (bool, float64) {
	// Treat sphere as plane, scale vector so that it is at same distance as
	// plane then check X and Y for collision
	v = v.Scale(s.Centre.Z / v.Z)
	// now check if in bounds
	distance := s.distanceFromCentre(v)

	return distance < s.Radius, distance
}

func (s Sphere) distanceFromCentre(v *Vector) float64 {
	return math.Sqrt(
		(v.X-s.Centre.X)*(v.X-s.Centre.X) +
			(v.Y-s.Centre.Y)*(v.Y-s.Centre.Y) +
			(v.Z-s.Centre.Z)*(v.Y-s.Centre.Z))
}
