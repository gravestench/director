package debug

import lua "github.com/yuin/gopher-lua"

// ExportToLua exports the component into the given table, using the given lua state machine.
func (*Debug) ExportToLua(_ *lua.LState, table *lua.LTable) *lua.LTable {
	// debug component contains no fields, it is merely a tag component

	return table
}
