package debug

import (
	"github.com/gravestench/akara"
	lua "github.com/yuin/gopher-lua"
)

const (
	componentName = "debug"
)

func (concrete *ComponentFactory) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	debugTable := state.NewTypeMetatable(componentName)

	state.SetField(debugTable, "add", concrete.luaAdd(state))
	state.SetField(debugTable, "get", concrete.luaGet(state))
	state.SetField(debugTable, "remove", concrete.luaRemove(state))

	state.SetField(table, componentName, debugTable)

	return table
}

func (concrete *ComponentFactory) luaAdd(state *lua.LState) *lua.LFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := akara.EID(state.CheckNumber(1))

		concrete.Add(e)

		L.Push(state.NewTable()) // intentionally blank, this is a tag component

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
		_, found := concrete.Get(akara.EID(id))

		truthy := lua.LFalse
		if !found {
			L.Push(lua.LNil)
			L.Push(truthy)
			return 2
		} else {
			truthy = lua.LTrue
		}

		// intentionally left blank
		// this component is just used for tagging an entity for debug purposes
		table := state.NewTable()

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
