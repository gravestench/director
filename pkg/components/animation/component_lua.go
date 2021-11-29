package animation

import lua "github.com/yuin/gopher-lua"

const (
	fieldFrame = "frame"
)

// ExportToLua exports the component into the given table, using the given lua state machine.
func (component *Animation) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	component.luaExportFrame(state, table)

	return table
}

func (component *Animation) luaExportFrame(state *lua.LState, table *lua.LTable) *lua.LTable {
	setterGetter := func(L *lua.LState) int {
		if L.GetTop() == 0 {
			L.Push(lua.LNumber(component.CurrentFrame))

			return 1
		}

		idx := int(L.CheckNumber(1))
		component.CurrentFrame = idx

		return 0
	}

	state.SetField(table, fieldFrame, state.NewFunction(setterGetter))

	return table
}
