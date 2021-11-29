package fill

import (
	"image/color"

	lua "github.com/yuin/gopher-lua"
)

const (
	fieldRGBA = "rgba"
)

// ExportToLua exports the component into the given table, using the given lua state machine.
func (component *Fill) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	component.luaExportFieldRGBA(state, table)

	return table
}

func (component *Fill) luaExportFieldRGBA(state *lua.LState, table *lua.LTable) {
	fnGetSet := func(L *lua.LState) int {
		if L.GetTop() == 4 {
			r := uint8(L.CheckNumber(1))
			g := uint8(L.CheckNumber(2))
			b := uint8(L.CheckNumber(3))
			a := uint8(L.CheckNumber(4))

			component.Color = color.RGBA{
				R: r,
				G: g,
				B: b,
				A: a,
			}

			return 0
		}

		r, g, b, a := component.RGBA()

		L.Push(lua.LNumber(r))
		L.Push(lua.LNumber(g))
		L.Push(lua.LNumber(b))
		L.Push(lua.LNumber(a))

		return 4
	}

	state.SetField(table, fieldRGBA, state.NewFunction(fnGetSet))
}
