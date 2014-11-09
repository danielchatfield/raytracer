package raytracer

import "testing"

func TestGetDirectionVector(t *testing.T) {
    camera := &Camera{}

    correct := Vector{X: 0, Y: 0, Z: 1}

    result := camera.getDirectionVector()

    if *result != correct {
        t.Errorf("Got %+v instead of %+v", *result, correct)
    }
}
