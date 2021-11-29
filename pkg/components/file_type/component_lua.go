package file_type

import (
	lua "github.com/yuin/gopher-lua"
)

const (
	fieldType = "type"
)

// ExportToLua exports the component into the given table, using the given lua state machine.
func (component *FileType) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	component.luaExportFieldType(state, table)

	return table
}

func (component *FileType) luaExportFieldType(state *lua.LState, table *lua.LTable) {
	fnGetSetType := func(L *lua.LState) int {
		if L.GetTop() == 1 {
			component.Type = L.CheckString(1)

			return 0
		}

		L.Push(lua.LString(component.Type))

		return 1
	}

	state.SetField(table, fieldType, state.NewFunction(fnGetSetType))
}
