package interactive

import (
	"fmt"

	"github.com/gravestench/director/pkg/systems/input/constants"

	components2 "github.com/gravestench/director/pkg/common/components"
	lua "github.com/yuin/gopher-lua"
)

// ExportToLua exports the component into the given table, using the given lua state machine.
func (component *Interactive) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	fnSetKey := func(L *lua.LState) int {
		if L.GetTop() < 1 {
			return 0
		}

		component.Vector.SetKey(constants.Key(L.CheckNumber(1)))

		return 0
	}

	fnSetMod := func(L *lua.LState) int {
		if L.GetTop() < 1 {
			return 0
		}

		component.Vector.SetModifier(constants.Modifier(L.CheckNumber(1)))

		return 0
	}

	fnSetMouseButton := func(L *lua.LState) int {
		if L.GetTop() < 1 {
			return 0
		}

		component.Vector.SetMouseButton(constants.MouseButton(L.CheckNumber(1)))

		return 0
	}

	state.SetField(table, "setKey", state.NewFunction(fnSetKey))
	state.SetField(table, "setMod", state.NewFunction(fnSetMod))
	state.SetField(table, "setMouse", state.NewFunction(fnSetMouseButton))

	hitBoxTable := state.NewTable()
	components2.LuaExportRectangle(component.Hitbox, state, hitBoxTable)

	callbackTable := state.NewTable()
	LuaExportInteractiveCallback(component, state, callbackTable)

	state.SetField(table, "hitbox", hitBoxTable)
	state.SetField(table, "callback", callbackTable)

	return table
}

func LuaExportInteractiveCallback(in *Interactive, state *lua.LState, table *lua.LTable) {
	fn := func(L *lua.LState) int {
		returnCount := 0

		switch L.GetTop() {
		case 0: // we are retrieving the function
			L.Push(state.NewFunction(func(L *lua.LState) int {
				in.Callback()

				return 0
			}))

			returnCount = 1
		case 1: // we are setting the function
			fnName := L.CheckString(1)
			in.Callback = func() bool {
				err := state.CallByParam(lua.P{
					Fn:      state.GetGlobal(fnName),
					NRet:    0,
					Protect: true,
				})

				if err != nil {
					fmt.Print(err)
				}

				return false // TODO handle propagation terminating callback
			}
		}

		return returnCount
	}

	state.NewFunction(fn)
}
