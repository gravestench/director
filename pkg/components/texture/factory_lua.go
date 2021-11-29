package texture

import lua "github.com/yuin/gopher-lua"

func (m *ComponentFactory) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	// not implemented, but dont want to panic here
	return table
}
