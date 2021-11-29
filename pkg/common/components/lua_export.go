package components

import lua "github.com/yuin/gopher-lua"

// LuaExport is an interface for things that can be made
// available inside of the lua environment.
type LuaExport interface {
	ExportToLua(*lua.LState, *lua.LTable) *lua.LTable
}
