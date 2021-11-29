package font

import lua "github.com/yuin/gopher-lua"

const (
	fieldFace = "face"
	fieldSize = "size"
)

// ExportToLua exports the component into the given table, using the given lua state machine.
func (component *Font) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	state.SetField(table, fieldFace, state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 0 {
			L.Push(lua.LString(component.Face))

			return 1
		}

		face := L.CheckString(1)
		component.Face = face

		return 0
	}))

	state.SetField(table, fieldSize, state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 0 {
			L.Push(lua.LNumber(component.Size))

			return 1
		}

		size := int(L.CheckNumber(1))
		component.Size = size

		return 0
	}))

	return table
}
