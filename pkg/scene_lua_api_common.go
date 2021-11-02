package pkg

import (
	"fmt"
	"image"

	"github.com/gravestench/mathlib"
	lua "github.com/yuin/gopher-lua"

	"github.com/gravestench/director/pkg/systems/input"
)

func (s *Scene) makeLuaSetterGetterVec3(vec3 *mathlib.Vector3) *lua.LFunction {
	setGetXYZ := func(L *lua.LState) int {
		if L.GetTop() == 3 {
			x := L.CheckNumber(1)
			y := L.CheckNumber(2)
			z := L.CheckNumber(3)

			vec3.Set(float64(x), float64(y), float64(z))

			return 0
		}

		x, y, z := vec3.XYZ()

		s.Lua.Push(lua.LNumber(x))
		s.Lua.Push(lua.LNumber(y))
		s.Lua.Push(lua.LNumber(z))

		return 3
	}

	return s.Lua.NewFunction(setGetXYZ)
}

func (s *Scene) makeLuaTableImageRectangle(r *image.Rectangle) *lua.LTable {
	table := s.Lua.NewTable()

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

	s.Lua.SetField(table, "position", s.Lua.NewFunction(getSetPosition))
	s.Lua.SetField(table, "size", s.Lua.NewFunction(getSetSize))

	return table
}

func (s *Scene) makeLuaInteractiveCallbackSetGet(in *input.Interactive) *lua.LFunction {
	fn := func(L *lua.LState) int {
		returnCount := 0

		switch L.GetTop() {
		case 0: // we are retrieving the function
			L.Push(s.Lua.NewFunction(func(L *lua.LState) int {
				in.Callback()

				return 0
			}))

			returnCount = 1
		case 1: // we are setting the function
			fnName := L.CheckString(1)
			in.Callback = func() bool {
				err := s.Lua.CallByParam(lua.P{
					Fn:      s.Lua.GetGlobal(fnName),
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

	return s.Lua.NewFunction(fn)
}
