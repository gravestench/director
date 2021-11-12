package pkg

import (
	lua "github.com/yuin/gopher-lua"

	"github.com/gravestench/director/pkg/systems/input"
)

const (
	luaInputConstantsName = "input"
)

func (s *SceneSystem) luaExportConstantsInput(mt *lua.LTable) {
	keyTable := s.Lua.NewTypeMetatable(luaInputConstantsName)

	keys := map[string]input.Key{
		"Key0":            input.Key0,
		"Key1":            input.Key1,
		"Key2":            input.Key2,
		"Key3":            input.Key3,
		"Key4":            input.Key4,
		"Key5":            input.Key5,
		"Key6":            input.Key6,
		"Key7":            input.Key7,
		"Key8":            input.Key8,
		"Key9":            input.Key9,
		"KeyA":            input.KeyA,
		"KeyB":            input.KeyB,
		"KeyC":            input.KeyC,
		"KeyD":            input.KeyD,
		"KeyE":            input.KeyE,
		"KeyF":            input.KeyF,
		"KeyG":            input.KeyG,
		"KeyH":            input.KeyH,
		"KeyI":            input.KeyI,
		"KeyJ":            input.KeyJ,
		"KeyK":            input.KeyK,
		"KeyL":            input.KeyL,
		"KeyM":            input.KeyM,
		"KeyN":            input.KeyN,
		"KeyO":            input.KeyO,
		"KeyP":            input.KeyP,
		"KeyQ":            input.KeyQ,
		"KeyR":            input.KeyR,
		"KeyS":            input.KeyS,
		"KeyT":            input.KeyT,
		"KeyU":            input.KeyU,
		"KeyV":            input.KeyV,
		"KeyW":            input.KeyW,
		"KeyX":            input.KeyX,
		"KeyY":            input.KeyY,
		"KeyZ":            input.KeyZ,
		"KeyApostrophe":   input.KeyApostrophe,
		"KeyBackslash":    input.KeyBackslash,
		"KeyBackspace":    input.KeyBackspace,
		"KeyCapsLock":     input.KeyCapsLock,
		"KeyComma":        input.KeyComma,
		"KeyDelete":       input.KeyDelete,
		"KeyDown":         input.KeyDown,
		"KeyEnd":          input.KeyEnd,
		"KeyEnter":        input.KeyEnter,
		"KeyEqual":        input.KeyEqual,
		"KeyEscape":       input.KeyEscape,
		"KeyF1":           input.KeyF1,
		"KeyF2":           input.KeyF2,
		"KeyF3":           input.KeyF3,
		"KeyF4":           input.KeyF4,
		"KeyF5":           input.KeyF5,
		"KeyF6":           input.KeyF6,
		"KeyF7":           input.KeyF7,
		"KeyF8":           input.KeyF8,
		"KeyF9":           input.KeyF9,
		"KeyF10":          input.KeyF10,
		"KeyF11":          input.KeyF11,
		"KeyF12":          input.KeyF12,
		"KeyGraveAccent":  input.KeyGraveAccent,
		"KeyHome":         input.KeyHome,
		"KeyInsert":       input.KeyInsert,
		"KeyKP0":          input.KeyKP0,
		"KeyKP1":          input.KeyKP1,
		"KeyKP2":          input.KeyKP2,
		"KeyKP3":          input.KeyKP3,
		"KeyKP4":          input.KeyKP4,
		"KeyKP5":          input.KeyKP5,
		"KeyKP6":          input.KeyKP6,
		"KeyKP7":          input.KeyKP7,
		"KeyKP8":          input.KeyKP8,
		"KeyKP9":          input.KeyKP9,
		"KeyKPAdd":        input.KeyKPAdd,
		"KeyKPDecimal":    input.KeyKPDecimal,
		"KeyKPDivide":     input.KeyKPDivide,
		"KeyKPEnter":      input.KeyKPEnter,
		"KeyKPEqual":      input.KeyKPEqual,
		"KeyKPMultiply":   input.KeyKPMultiply,
		"KeyKPSubtract":   input.KeyKPSubtract,
		"KeyLeft":         input.KeyLeft,
		"KeyLeftBracket":  input.KeyLeftBracket,
		"KeyMenu":         input.KeyMenu,
		"KeyMinus":        input.KeyMinus,
		"KeyNumLock":      input.KeyNumLock,
		"KeyPageDown":     input.KeyPageDown,
		"KeyPageUp":       input.KeyPageUp,
		"KeyPause":        input.KeyPause,
		"KeyPeriod":       input.KeyPeriod,
		"KeyPrintScreen":  input.KeyPrintScreen,
		"KeyRight":        input.KeyRight,
		"KeyRightBracket": input.KeyRightBracket,
		"KeyScrollLock":   input.KeyScrollLock,
		"KeySemicolon":    input.KeySemicolon,
		"KeySlash":        input.KeySlash,
		"KeySpace":        input.KeySpace,
		"KeyTab":          input.KeyTab,
		"KeyUp":           input.KeyUp,

		"MouseButtonLeft":   input.MouseButtonLeft,
		"MouseButtonRight":  input.MouseButtonRight,
		"MouseButtonMiddle": input.MouseButtonMiddle,

		"ModAltLeft":      input.ModAltLeft,
		"ModAltRight":     input.ModAltRight,
		"ModControlLeft":  input.ModControlLeft,
		"ModControlRight": input.ModControlRight,
		"ModShiftLeft":    input.ModShiftLeft,
		"ModShiftRight":   input.ModShiftRight,
	}

	for k, v := range keys {
		s.Lua.SetField(keyTable, k, lua.LNumber(v))
	}

	s.Lua.SetField(mt, luaInputConstantsName, keyTable)
}
