package constants

import rl "github.com/gen2brain/raylib-go/raylib"

// Modifier represents a keyboard modifier key
type Modifier = int32

// Modifiers
const (
	ModAltLeft      = Modifier(rl.KeyLeftAlt)
	ModAltRight     = Modifier(rl.KeyRightAlt)
	ModControlLeft  = Modifier(rl.KeyLeftControl)
	ModControlRight = Modifier(rl.KeyRightControl)
	ModShiftLeft    = Modifier(rl.KeyLeftShift)
	ModShiftRight   = Modifier(rl.KeyRightShift)
)
