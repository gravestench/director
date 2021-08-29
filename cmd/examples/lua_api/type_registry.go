package main

import (
	"github.com/gravestench/director/cmd/examples/lua_api/common"
	"github.com/gravestench/director/cmd/examples/lua_api/vector3"
	lua "github.com/yuin/gopher-lua"
)

var luaTypes = []common.LuaTypeExport{
	vector3.LuaTypeExport,
}

// registerType takes a LuaTypeExport
func registerType(L *lua.LState, luaTypeExport common.LuaTypeExport) {
	typeMetatable := L.NewTypeMetatable(luaTypeExport.Name)
	L.SetGlobal(luaTypeExport.Name, typeMetatable)

	// static attributes
	L.SetField(typeMetatable, "new", L.NewFunction(luaTypeExport.ConstructorFunc))

	// methods
	L.SetField(typeMetatable, "__index", L.SetFuncs(L.NewTable(), luaTypeExport.Methods))
}