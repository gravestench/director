package file_load_request

import (
	"github.com/gravestench/akara"
	lua "github.com/yuin/gopher-lua"
)

const (
	componentName = "fileLoadRequest"
)

// ExportToLua exports the component factory to the target table using the given lua state machine
func (f *ComponentFactory) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	factoryTable := state.NewTypeMetatable(componentName)

	state.SetField(factoryTable, "add", f.luaAdd(state))
	state.SetField(factoryTable, "get", f.luaGet(state))
	state.SetField(factoryTable, "remove", f.luaRemove(state))

	state.SetField(table, componentName, factoryTable)

	return table
}

func (f *ComponentFactory) luaAdd(state *lua.LState) *lua.LFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := akara.EID(state.CheckNumber(1))

		fileLoadRequest := f.Add(e)
		L.Push(fileLoadRequest.ExportToLua(state, state.NewTable()))
		return 1
	}

	return state.NewFunction(fn)
}

func (f *ComponentFactory) luaGet(state *lua.LState) *lua.LFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		id := L.CheckNumber(1)
		fileLoadRequest, found := f.Get(akara.EID(id))

		truthy := lua.LFalse
		if !found {
			L.Push(lua.LNil)
			L.Push(truthy)
			return 2
		} else {
			truthy = lua.LTrue
		}

		table := fileLoadRequest.ExportToLua(state, state.NewTable())

		L.SetMetatable(table, L.GetTypeMetatable(componentName))

		L.Push(table)
		L.Push(truthy)

		return 2
	}

	return state.NewFunction(fn)
}

func (f *ComponentFactory) luaRemove(state *lua.LState) *lua.LFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := akara.EID(state.CheckNumber(1))

		f.Remove(e)

		return 0
	}

	return state.NewFunction(fn)
}
