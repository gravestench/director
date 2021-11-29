package viewport

import lua "github.com/yuin/gopher-lua"

// ExportToLua exports the component into the given table, using the given lua state machine.
func (concrete *ComponentFactory) ExportToLua(_ *lua.LState, table *lua.LTable) *lua.LTable {
	// not implemented, but dont want to panic here
	return table
}
