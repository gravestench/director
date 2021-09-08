package scene

import (
	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/mathlib"
	lua "github.com/yuin/gopher-lua"
)

func (s *Scene) luaCheckEID() *common.Entity {
	ud := s.Lua.CheckUserData(1)
	if v, ok := ud.Value.(*common.Entity); ok {
		return v
	}

	s.Lua.ArgError(1, "EID expected")
	return nil
}

func (s *Scene) makeLuaTableVec3(vec3 *mathlib.Vector3) *lua.LFunction {
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

func (s *Scene) makeLuaTableVec2(vec2 *mathlib.Vector2) *lua.LFunction {
	setGetXYZ := func(L *lua.LState) int {
		if L.GetTop() == 2 {
			x := L.CheckNumber(1)
			y := L.CheckNumber(2)

			vec2.Set(float64(x), float64(y))

			return 0
		}

		x, y := vec2.XY()

		s.Lua.Push(lua.LNumber(x))
		s.Lua.Push(lua.LNumber(y))

		return 3
	}

	return s.Lua.NewFunction(setGetXYZ)
}