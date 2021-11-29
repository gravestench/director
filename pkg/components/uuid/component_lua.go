package uuid

import (
	lua "github.com/yuin/gopher-lua"
)

const (
	fieldString = "string"
)

// ExportToLua exports the component into the given table, using the given lua state machine.
func (component *UUID) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	component.luaExportFieldString(state, table)

	return table
}

func (component *UUID) luaExportFieldString(state *lua.LState, table *lua.LTable) {
	fnSet := state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() != 0 {

			return 0
		}

		L.Push(lua.LString(component.String()))

		return 1
	})

	state.SetField(table, fieldString, fnSet)
}
