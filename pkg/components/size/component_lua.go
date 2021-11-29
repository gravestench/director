package size

import lua "github.com/yuin/gopher-lua"

const (
	fieldSize = "size"
)

// ExportToLua exports the component into the given table, using the given lua state machine.
func (component *Size) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	fnSetGet := state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() >= 2 {
			return 0
		}

		w, h := component.Max.X, component.Max.Y
		L.Push(lua.LNumber(w))
		L.Push(lua.LNumber(h))

		return 2
	})

	state.SetField(table, fieldSize, fnSetGet)

	return table
}
