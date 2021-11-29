package camera

import lua "github.com/yuin/gopher-lua"

const (
	fieldRotation = "rotation"
	fieldZoom     = "zoom"
)

// ExportToLua exports the component into the given table, using the given lua state machine.
func (component *Camera) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	component.luaExportRotation(state, table)
	component.luaExportZoom(state, table)

	return table
}

func (component *Camera) luaExportRotation(state *lua.LState, table *lua.LTable) {
	fnSetGet := func(L *lua.LState) int {
		if L.GetTop() == 0 {
			L.Push(lua.LNumber(component.Rotation))

			return 1
		}

		r := float32(L.CheckNumber(1))
		component.Rotation = r

		return 0
	}

	state.SetField(table, fieldRotation, state.NewFunction(fnSetGet))
}

func (component *Camera) luaExportZoom(state *lua.LState, table *lua.LTable) {
	fnSetGet := func(L *lua.LState) int {
		if L.GetTop() == 0 {
			L.Push(lua.LNumber(component.Zoom))

			return 1
		}

		z := float32(L.CheckNumber(1))
		component.Zoom = z

		return 0
	}

	state.SetField(table, fieldZoom, state.NewFunction(fnSetGet))
}
