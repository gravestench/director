package common

import lua "github.com/yuin/gopher-lua"

type LuaModuleExport = map[string]lua.LGFunction

func LoadLuaModule(L *lua.LState, exports LuaModuleExport) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), exports)
	// register other stuff
	L.SetField(mod, "name", lua.LString("value"))

	// returns the module
	L.Push(mod)
	return 1
}

//var exports = map[string]lua.LGFunction{
//	"myfunc": myfunc,
//}
//
//func myfunc(L *lua.LState) int {
//	return 0
//}
