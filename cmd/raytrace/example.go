package main

import . "github.com/danielchatfield/raytracer"

func main() {
	camera := NewCamera(500, 400)
	sphere := Sphere{
		Object: Object{
			Ambient:  0.9,
			Diffuse:  0.9,
			Specular: 0.9,
		},
		Centre: Vector{
			X: 0,
			Y: 0,
			Z: 50,
		},
		Radius: 5,
	}
	objects := []ObjectInterface{sphere}
	scene := Scene{
		Objects: objects,
	}
	renderer := NewRenderer(camera, &scene)

	renderer.Render()
}
