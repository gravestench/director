package input

import rl "github.com/gen2brain/raylib-go/raylib"

// MouseButton represents a button on a mouse
type MouseButton = int32

// MouseButtons
const (
	MouseButtonLeft   = MouseButton(rl.MouseLeftButton)
	MouseButtonRight  = MouseButton(rl.MouseRightButton)
	MouseButtonMiddle = MouseButton(rl.MouseMiddleButton)
)
