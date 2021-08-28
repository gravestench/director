package rectangle

import "image"

// Rectangle is a Go representation of a dog.
type Rectangle struct {
	image.Rectangle
}

// NewSquare creates a new Rectangle object.
func NewSquare(name string) Rectangle {
	return Rectangle{}
}
