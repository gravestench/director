package stroke

import (
	"image/color"

	lua "github.com/yuin/gopher-lua"
)

const (
	fieldRGBA      = "rgba"
	fieldLineWidth = "lineWidth"
)

// ExportToLua exports the component into the given table, using the given lua state machine.
func (component *Stroke) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	component.luaExportFieldRGBA(state, table)
	component.luaExportFieldLineWidth(state, table)

	return table
}

func (component *Stroke) luaExportFieldRGBA(state *lua.LState, table *lua.LTable) {
	fnGetSet := state.NewFunction(func(L *lua.LState) int {
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
	})

	state.SetField(table, fieldRGBA, fnGetSet)
}

func (component *Stroke) luaExportFieldLineWidth(state *lua.LState, table *lua.LTable) {
	fnGetSet := state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 1 {
			lw := int(L.CheckNumber(1))

			if lw < 0 {
				lw *= -1
			}

			component.LineWidth = lw

			return 0
		}

		L.Push(lua.LNumber(component.LineWidth))

		return 1
	})

	state.SetField(table, fieldRGBA, fnGetSet)
}
