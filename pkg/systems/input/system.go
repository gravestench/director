package input

import (
	"image"

	"github.com/gravestench/director/pkg/systems/input/vector"

	"github.com/gravestench/director/pkg/systems/input/constants"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/components/interactive"
	"github.com/gravestench/mathlib"
)

// static check that System implements the System interface
var _ akara.System = &System{}

// System is responsible for handling interactive entities
type System struct {
	akara.BaseSystem
	interactives  *akara.Subscription
	InputState    *vector.Vector
	MousePosition mathlib.Vector2
	Components    struct {
		Interactive interactive.ComponentFactory
	}
}

func (m *System) Name() string {
	return "Input"
}

func (m *System) IsInitialized() bool {
	return m.InputState != nil
}

// Init initializes the system with the given world, injecting the necessary components
func (m *System) Init(_ *akara.World) {
	m.setupFactories()
	m.setupSubscriptions()

	m.InputState = vector.NewInputVector()
}

func (m *System) setupFactories() {
	m.InjectComponent(&interactive.Interactive{}, &m.Components.Interactive.ComponentFactory)
}

func (m *System) setupSubscriptions() {
	interactives := m.NewComponentFilter().
		Require(&interactive.Interactive{}).
		Build()

	m.interactives = m.AddSubscription(interactives)
}

// Update will iterate over interactive entities
func (m *System) Update() {
	m.updateInputState()

	for _, id := range m.interactives.GetEntities() {
		preventPropagation := m.applyInputState(id)
		if preventPropagation {
			break
		}
	}
}

func (m *System) updateInputState() {
	m.InputState.Clear()

	var keysToCheck = []constants.Key{
		constants.Key0, constants.Key1, constants.Key2, constants.Key3, constants.Key4, constants.Key5, constants.Key6,
		constants.Key7, constants.Key8, constants.Key9, constants.KeyA, constants.KeyB, constants.KeyC, constants.KeyD,
		constants.KeyE, constants.KeyF, constants.KeyG, constants.KeyH, constants.KeyI, constants.KeyJ, constants.KeyK,
		constants.KeyL, constants.KeyM, constants.KeyN, constants.KeyO, constants.KeyP, constants.KeyQ, constants.KeyR,
		constants.KeyS, constants.KeyT, constants.KeyU, constants.KeyV, constants.KeyW, constants.KeyX, constants.KeyY,
		constants.KeyZ, constants.KeyApostrophe, constants.KeyBackslash, constants.KeyBackspace,
		constants.KeyCapsLock, constants.KeyComma, constants.KeyDelete, constants.KeyDown,
		constants.KeyEnd, constants.KeyEnter, constants.KeyEqual, constants.KeyEscape,
		constants.KeyF1, constants.KeyF2, constants.KeyF3, constants.KeyF4, constants.KeyF5, constants.KeyF6,
		constants.KeyF7, constants.KeyF8, constants.KeyF9, constants.KeyF10, constants.KeyF11, constants.KeyF12,
		constants.KeyGraveAccent, constants.KeyHome, constants.KeyInsert, constants.KeyKP0,
		constants.KeyKP1, constants.KeyKP2, constants.KeyKP3, constants.KeyKP4, constants.KeyKP5,
		constants.KeyKP6, constants.KeyKP7, constants.KeyKP8, constants.KeyKP9,
		constants.KeyKPAdd, constants.KeyKPDecimal, constants.KeyKPDivide, constants.KeyKPEnter,
		constants.KeyKPEqual, constants.KeyKPMultiply, constants.KeyKPSubtract, constants.KeyLeft,
		constants.KeyLeftBracket, constants.KeyMenu, constants.KeyMinus, constants.KeyNumLock,
		constants.KeyPageDown, constants.KeyPageUp, constants.KeyPause, constants.KeyPeriod,
		constants.KeyPrintScreen, constants.KeyRight, constants.KeyRightBracket,
		constants.KeyScrollLock, constants.KeySemicolon, constants.KeySlash,
		constants.KeySpace, constants.KeyTab, constants.KeyUp,
	}

	var modifiersToCheck = []constants.Modifier{
		constants.ModAltLeft, constants.ModAltRight,
		constants.ModControlLeft, constants.ModControlRight,
		constants.ModShiftLeft, constants.ModShiftRight,
	}

	var buttonsToCheck = []constants.MouseButton{
		constants.MouseButtonLeft, constants.MouseButtonMiddle, constants.MouseButtonRight,
	}

	for _, key := range keysToCheck {
		truth := rl.IsKeyDown(key)
		m.InputState.KeyVector.Set(int(key), truth)
	}

	for _, mod := range modifiersToCheck {
		truth := rl.IsKeyDown(mod)
		m.InputState.ModifierVector.Set(int(mod), truth)
	}

	for _, btn := range buttonsToCheck {
		truth := rl.IsMouseButtonDown(btn)
		m.InputState.MouseButtonVector.Set(int(btn), truth)
	}

	mousePos := rl.GetMousePosition()
	m.MousePosition.Set(float64(mousePos.X), float64(mousePos.Y))
}

func (m *System) applyInputState(id akara.EID) (preventPropagation bool) {
	i, _ := m.Components.Interactive.Get(id)

	// check if this Interactive specified a particular cursor position that the input must occur in
	if i.Hitbox != nil {
		if !contains(i.Hitbox, int(m.MousePosition.X), int(m.MousePosition.Y)) {
			i.UsedRecently = false
			return false
		}
	}

	// verify that the current InputState matches the state specified in the Vector
	if !i.Enabled || !m.InputState.Contains(i.Vector) {
		i.UsedRecently = false
		return false
	}

	// if this Interactive component is in rapid fire mode, we don't need this debouncing logic
	if !i.RapidFire {
		// if the input state hasn't changed since the last time we ran the callback, don't run it again.
		// This resets when the input state changes.
		if i.UsedRecently {
			return false
		}

		// we've verified that the current input state matches the vector.
		// mark the callback as having been run so we don't run it again until the input state changes.
		i.UsedRecently = true
	}

	return i.Callback()
}

func contains(r *image.Rectangle, x, y int) bool {
	return (r.Min.X < x && r.Max.X > x) && (r.Min.Y < y && r.Max.Y > y)
}
