package raytracer

import (
	"image"
	"image/png"
	"math"
	"os"
)

// Camera represents the scene camera.
type Camera struct {
	Width  int
	Height int
	Image  *image.RGBA
}

// Image is a 2d slice representing the pixels
type Image [][]Color

// Color represents a color in RGB
type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

func NewCamera(width int, height int) *Camera {
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))

	return &Camera{
		Width:  width,
		Height: height,
		Image:  rgba,
	}
}

func (c *Camera) getFieldOfView() float64 {
	return math.Pi / 3
}

func (c *Camera) getDirectionVector() *Vector {
	return &Vector{
		X: 0,
		Y: 0,
		Z: 1,
	}
}

func (c *Camera) getViewDistance() float64 {
	return float64(c.Width) / (2.0 * math.Tan(c.getFieldOfView()/2.0))
}

func (c *Camera) Snap() {
	dest, _ := os.Create("./dest.png")
	defer dest.Close()

	png.Encode(dest, c.Image)
}
