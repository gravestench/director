package text

import lua "github.com/yuin/gopher-lua"

const (
	fieldString = "string"
)

// ExportToLua exports the component into the given table, using the given lua state machine.
func (component *Text) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	component.luaExportFieldString(state, table)

	return table
}

func (component *Text) luaExportFieldString(state *lua.LState, table *lua.LTable) *lua.LTable {
	fnSetGet := state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 1 {
			val := L.CheckString(1)

			component.String = val

			return 0
		}

		L.Push(lua.LString(component.String))

		return 1
	})

	state.SetField(table, fieldString, fnSetGet)

	return table
}
