package opacity

import lua "github.com/yuin/gopher-lua"

const (
	fieldValue = "value"
)

// ExportToLua exports the component into the given table, using the given lua state machine.
func (component *Opacity) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	component.luaExportValue(state, table)

	return table
}

func (component *Opacity) luaExportValue(state *lua.LState, table *lua.LTable) {
	fnSetGet := state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 1 {
			val := float64(L.CheckNumber(1))

			component.Value = val

			return 0
		}

		L.Push(lua.LNumber(component.Value))

		return 1
	})

	state.SetField(table, fieldValue, fnSetGet)
}
