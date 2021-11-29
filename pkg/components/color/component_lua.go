package color

import (
	"image/color"

	lua "github.com/yuin/gopher-lua"
)

const (
	fieldRGBA = "rgba"
)

// ExportToLua exports the component into the given table, using the given lua state machine.
func (component *Color) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	component.luaExportRGBA(state, table)

	return table
}

func (component *Color) luaExportRGBA(state *lua.LState, table *lua.LTable) {
	fnSetGet := func(L *lua.LState) int {
		if L.GetTop() == 4 {
			r := uint8(L.CheckNumber(1))
			g := uint8(L.CheckNumber(1))
			b := uint8(L.CheckNumber(1))
			a := uint8(L.CheckNumber(1))

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

	state.SetField(table, fieldRGBA, state.NewFunction(fnSetGet))
}
