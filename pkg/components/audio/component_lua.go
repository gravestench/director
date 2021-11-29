package audio

import lua "github.com/yuin/gopher-lua"

// ExportToLua exports the factory to a lua table using the given lua state machine and lua table
func (a *Audio) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	// not implemented, but dont want to panic here
	return table
}
