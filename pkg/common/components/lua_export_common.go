package components

import (
	"image"

	"github.com/gravestench/mathlib"
	lua "github.com/yuin/gopher-lua"
)

func LuaExportVector3(vec3 *mathlib.Vector3, state *lua.LState, table *lua.LTable) *lua.LTable {
	const (
		fieldXYZ = "xyz"
	)

	setGetXYZ := func(L *lua.LState) int {
		if L.GetTop() == 3 {
			x := L.CheckNumber(1)
			y := L.CheckNumber(2)
			z := L.CheckNumber(3)

			vec3.Set(float64(x), float64(y), float64(z))

			return 0
		}

		x, y, z := vec3.XYZ()

		state.Push(lua.LNumber(x))
		state.Push(lua.LNumber(y))
		state.Push(lua.LNumber(z))

		return 3
	}

	state.SetField(table, fieldXYZ, state.NewFunction(setGetXYZ))

	return table
}

func LuaExportRectangle(r *image.Rectangle, state *lua.LState, table *lua.LTable) *lua.LTable {
	getSetPosition := func(L *lua.LState) int {
		returnCount := 0

		switch L.GetTop() {
		case 0:
			L.Push(lua.LNumber(r.Min.X))
			L.Push(lua.LNumber(r.Min.Y))

			returnCount = 2
		case 2:
			r.Min.X = int(L.CheckNumber(1))
			r.Min.Y = int(L.CheckNumber(2))
		}

		return returnCount
	}

	getSetSize := func(L *lua.LState) int {
		returnCount := 0

		switch L.GetTop() {
		case 0:
			L.Push(lua.LNumber(r.Dx()))
			L.Push(lua.LNumber(r.Dy()))

			returnCount = 2
		case 2:
			w := int(L.CheckNumber(1))
			h := int(L.CheckNumber(2))

			r.Max.X = r.Min.X + w
			r.Max.Y = r.Min.Y + h
		}

		return returnCount
	}

	state.SetField(table, "position", state.NewFunction(getSetPosition))
	state.SetField(table, "size", state.NewFunction(getSetSize))

	return table
}
