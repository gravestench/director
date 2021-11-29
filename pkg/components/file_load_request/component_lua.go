package file_load_request

import lua "github.com/yuin/gopher-lua"

const (
	fieldPath = "path"
)

// ExportToLua exports the component into the given table, using the given lua state machine.
func (component *FileLoadRequest) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	component.luaExportFieldPath(state, table)

	return table
}

func (component *FileLoadRequest) luaExportFieldPath(state *lua.LState, table *lua.LTable) {
	fnGetSet := func(L *lua.LState) int {
		if L.GetTop() == 1 {
			uri := L.CheckString(1)

			component.Path = uri

			return 0
		}

		L.Push(lua.LString(component.Path))

		return 1
	}

	state.SetField(table, fieldPath, state.NewFunction(fnGetSet))
}
