package render_order

import lua "github.com/yuin/gopher-lua"

// ExportToLua exports the component into the given table, using the given lua state machine.
func (component *RenderOrder) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	state.SetField(table, "value", state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 1 {
			val := int(L.CheckNumber(1))

			component.Value = val

			return 0
		}

		L.Push(lua.LNumber(component.Value))

		return 1
	}))

	return table
}
