package raytracer

import "testing"

func TestDistanceFromCentre(t *testing.T) {
	sphere := Sphere{
		Centre: Vector{
			X: 1,
			Y: 2,
			Z: 10,
		},
		Radius: 4,
	}

	v := &Vector{
		X: 4,
		Y: -2,
		Z: 10,
	}

	correct := 5.0

	result := sphere.distanceFromCentre(v)

	if result != correct {
		t.Errorf("Got %+v instead of %+v", result, correct)
	}
}
