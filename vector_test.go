package raytracer

import "testing"

func TestAdd(t *testing.T) {
	v1 := &Vector{X: 1, Y: 2, Z: 3}
	v2 := &Vector{X: 1, Y: 2.5, Z: -1}

	result := Vector{X: 2, Y: 4.5, Z: 2}

	rv := v1.Add(v2)

	if *rv != result {
		t.Errorf("Got %+v instead of %+v", *rv, result)
	}
}

func TestSubtract(t *testing.T) {
	v1 := &Vector{X: 1, Y: 2, Z: 3}
	v2 := &Vector{X: 1, Y: 2.5, Z: -1}

	result := Vector{X: 0, Y: -0.5, Z: 4}

	rv := v1.Subtract(v2)

	if *rv != result {
		t.Errorf("Got %+v instead of %+v", *rv, result)
	}
}

func TestScale(t *testing.T) {
	v := &Vector{X: 1, Y: 2, Z: 3}

	result := Vector{X: -1.5, Y: -3, Z: -4.5}

	rv := v.Scale(-1.5)

	if *rv != result {
		t.Errorf("Got %+v instead of %+v", *rv, result)
	}
}

func TestDotProduct(t *testing.T) {
	v1 := &Vector{X: 1, Y: 2, Z: 3}
	v2 := &Vector{X: 3.5, Y: -1, Z: 1}

	result := 4.5

	rv := v1.DotProduct(v2)

	if rv != result {
		t.Errorf("Got %v instead of %v", rv, result)
	}
}
