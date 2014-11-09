package raytracer

import (
	"image/color"

	. "github.com/tj/go-debug"
)

var debug = Debug("renderer")

// Renderer renders the scene
type Renderer struct {
	camera *Camera
	scene  *Scene
}

// NewRenderer creates a renderer struct
func NewRenderer(camera *Camera, scene *Scene) *Renderer {
	return &Renderer{
		camera: camera,
		scene:  scene,
	}
}

// Render renders the scene
func (r *Renderer) Render() {

	for x := 0; x < r.camera.Width; x++ {
		for y := 0; y < r.camera.Height; y++ {
			r.renderPixel(x, y)
		}
	}

	r.camera.Snap()
}

func (r *Renderer) renderPixel(x int, y int) {
	z := r.camera.getViewDistance()
	halfWidth := r.camera.Width / 2.0
	halfHeight := r.camera.Height / 2.0

	direction := &Vector{
		X: float64(x) - float64(halfWidth) + 0.5,
		Y: float64(y) - float64(halfHeight) + 0.5,
		Z: z,
	}
	direction = direction.UnitVector()

	v := &Ray{
		Direction: *direction,
	}

	minDistance := -1.0
	var collisionObject Sphere
	var collisionPosition *Vector

	for _, object := range r.scene.Objects {
		doesIntercept, position := object.getInterception(v)

		if doesIntercept {
			if minDistance == -1 || position.Length() < minDistance {
				minDistance = position.Length()
				collisionObject = object.(Sphere)
				collisionPosition = position
			}
		}
	}

	color := color.NRGBA{0, 0, 0, 255}

	if minDistance != -1 {
		color = collisionObject.getColor(collisionPosition)
	}

	r.camera.Image.Set(x, y, color)

	debug("Rendering pixel (%v,%v) with color %V", x, y, color)
}
