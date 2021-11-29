package interactive

import (
	"github.com/gravestench/akara"
	lua "github.com/yuin/gopher-lua"
)

const (
	componentName = "interactive"
)

// ExportToLua exports the component into the given table, using the given lua state machine.
func (concrete *ComponentFactory) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	factoryTable := state.NewTable()

	fnAdd := state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := akara.EID(state.CheckNumber(1))

		interactive := concrete.Add(e)
		L.Push(interactive.ExportToLua(state, state.NewTable()))
		return 1
	})

	fnGet := state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		id := L.CheckNumber(1)
		interactive, found := concrete.Get(akara.EID(id))

		truthy := lua.LFalse
		if !found {
			L.Push(lua.LNil)
			L.Push(truthy)
			return 2
		} else {
			truthy = lua.LTrue
		}

		table := interactive.ExportToLua(state, state.NewTable())

		L.SetMetatable(table, L.GetTypeMetatable(componentName))

		L.Push(table)
		L.Push(truthy)

		return 2
	})

	fnRemove := state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := akara.EID(state.CheckNumber(1))

		concrete.Remove(e)

		return 1
	})

	state.SetField(factoryTable, "add", fnAdd)
	state.SetField(factoryTable, "get", fnGet)
	state.SetField(factoryTable, "remove", fnRemove)

	state.SetField(table, componentName, factoryTable)

	return table
}
