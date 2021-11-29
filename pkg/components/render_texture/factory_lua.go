package render_texture

import lua "github.com/yuin/gopher-lua"

func (concrete *ComponentFactory) ExportToLua(_ *lua.LState, table *lua.LTable) *lua.LTable {
	// not implemented, but dont want to panic here
	return table
}
