package file_load_response

import (
	"github.com/gravestench/akara"
	lua "github.com/yuin/gopher-lua"
)

const componentName = "fileLoadResponse"

// ExportToLua exports the factory to a lua table using the given lua state machine and lua table
func (concrete *ComponentFactory) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	factoryTable := state.NewTypeMetatable(componentName)

	state.SetField(factoryTable, "add", concrete.luaAdd(state))
	state.SetField(factoryTable, "get", concrete.luaGet(state))
	state.SetField(factoryTable, "remove", concrete.luaRemove(state))

	state.SetField(table, componentName, factoryTable)

	return table
}

func (concrete *ComponentFactory) luaAdd(state *lua.LState) *lua.LFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := akara.EID(state.CheckNumber(1))

		fileLoadResponse := concrete.Add(e)
		L.Push(fileLoadResponse.ExportToLua(state, state.NewTable()))
		return 1
	}

	return state.NewFunction(fn)
}

func (concrete *ComponentFactory) luaGet(state *lua.LState) *lua.LFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		id := L.CheckNumber(1)
		fileLoadResponse, found := concrete.Get(akara.EID(id))

		truthy := lua.LFalse
		if !found {
			L.Push(lua.LNil)
			L.Push(truthy)
			return 2
		} else {
			truthy = lua.LTrue
		}

		table := fileLoadResponse.ExportToLua(state, state.NewTable())

		L.SetMetatable(table, L.GetTypeMetatable(componentName))

		L.Push(table)
		L.Push(truthy)

		return 2
	}

	return state.NewFunction(fn)
}

func (concrete *ComponentFactory) luaRemove(state *lua.LState) *lua.LFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := akara.EID(state.CheckNumber(1))

		concrete.Remove(e)

		return 0
	}

	return state.NewFunction(fn)
}
