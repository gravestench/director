package input

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	"github.com/gravestench/mathlib"
	"image"
)

// static check that System implements the System interface
var _ akara.System = &System{}

// System is responsible for handling interactive entities
type System struct {
	akara.BaseSystem
	interactives  *akara.Subscription
	InputState    *Vector
	MousePosition mathlib.Vector2
	Components    struct {
		Interactive InteractiveFactory
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

	m.InputState = NewInputVector()
}

func (m *System) setupFactories() {
	m.InjectComponent(&Interactive{}, &m.Components.Interactive.ComponentFactory)
}

func (m *System) setupSubscriptions() {
	interactives := m.NewComponentFilter().
		Require(&Interactive{}).
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

	var keysToCheck = []Key{
		Key0, Key1, Key2, Key3, Key4, Key5, Key6,
		Key7, Key8, Key9, KeyA, KeyB, KeyC, KeyD,
		KeyE, KeyF, KeyG, KeyH, KeyI, KeyJ, KeyK,
		KeyL, KeyM, KeyN, KeyO, KeyP, KeyQ, KeyR,
		KeyS, KeyT, KeyU, KeyV, KeyW, KeyX, KeyY,
		KeyZ, KeyApostrophe, KeyBackslash, KeyBackspace,
		KeyCapsLock, KeyComma, KeyDelete, KeyDown,
		KeyEnd, KeyEnter, KeyEqual, KeyEscape,
		KeyF1, KeyF2, KeyF3, KeyF4, KeyF5, KeyF6,
		KeyF7, KeyF8, KeyF9, KeyF10, KeyF11, KeyF12,
		KeyGraveAccent, KeyHome, KeyInsert, KeyKP0,
		KeyKP1, KeyKP2, KeyKP3, KeyKP4, KeyKP5,
		KeyKP6, KeyKP7, KeyKP8, KeyKP9,
		KeyKPAdd, KeyKPDecimal, KeyKPDivide, KeyKPEnter,
		KeyKPEqual, KeyKPMultiply, KeyKPSubtract, KeyLeft,
		KeyLeftBracket, KeyMenu, KeyMinus, KeyNumLock,
		KeyPageDown, KeyPageUp, KeyPause, KeyPeriod,
		KeyPrintScreen, KeyRight, KeyRightBracket,
		KeyScrollLock, KeySemicolon, KeySlash,
		KeySpace, KeyTab, KeyUp,
	}

	var modifiersToCheck = []Modifier{
		ModAltLeft, ModAltRight,
		ModControlLeft, ModControlRight,
		ModShiftLeft, ModShiftRight,
	}

	var buttonsToCheck = []MouseButton{
		MouseButtonLeft, MouseButtonMiddle, MouseButtonRight,
	}

	for _, key := range keysToCheck {
		truth := rl.IsKeyPressed(key)
		m.InputState.KeyVector.Set(int(key), truth)
	}

	for _, mod := range modifiersToCheck {
		truth := rl.IsKeyPressed(mod)
		m.InputState.ModifierVector.Set(int(mod), truth)
	}

	for _, btn := range buttonsToCheck {
		truth := rl.IsMouseButtonPressed(btn)
		m.InputState.MouseButtonVector.Set(int(btn), truth)
	}

	mousePos := rl.GetMousePosition()
	m.MousePosition.Set(float64(mousePos.X), float64(mousePos.Y))
}

func (m *System) applyInputState(id akara.EID) (preventPropagation bool) {
	i, found := m.Components.Interactive.Get(id)
	if !found {
		return false
	}

	// verify that the current InputState matches the state specified in the Vector
	if !i.Enabled || !m.InputState.Contains(i.Vector) {
		return false
	}

	// check if this Interactive specified a particular cursor position that the input must occur in
	if i.Hitbox != nil {
		if !contains(i.Hitbox, int(m.MousePosition.X), int(m.MousePosition.Y)) {
			return false
		}
	}

	return i.Callback()
}

func contains(r *image.Rectangle, x, y int) bool {
	return (r.Min.X < x && r.Max.X > x) && (r.Min.Y < y && r.Max.Y > y)
}
