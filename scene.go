package raytracer

// Scene represents the world being rendered
type Scene struct {
	Objects []ObjectInterface
	Lights  []Light
}
