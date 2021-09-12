package common

import (
	lua "github.com/yuin/gopher-lua"
)

// LuaTypeExport is a collection of all the information we
// need about a Go type in order to export it for use in Lua scripts.
type LuaTypeExport struct {
	Name            string
	ConstructorFunc lua.LGFunction
	Methods         map[string]lua.LGFunction
}

// RegisterLuaType takes a LuaTypeExport and uses it to add a new global to the Lua state machine
func RegisterLuaType(L *lua.LState, luaTypeExport LuaTypeExport) {
	typeMetatable := L.NewTypeMetatable(luaTypeExport.Name)
	L.SetGlobal(luaTypeExport.Name, typeMetatable)

	// static attributes
	L.SetField(typeMetatable, "new", L.NewFunction(luaTypeExport.ConstructorFunc))

	// methods
	L.SetField(typeMetatable, "__index", L.SetFuncs(L.NewTable(), luaTypeExport.Methods))
}
