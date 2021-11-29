package texture

import lua "github.com/yuin/gopher-lua"

// ExportToLua exports the component into the given table, using the given lua state machine.
func (d *Texture2D) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	// not implemented, but dont want to panic here
	return table
}
