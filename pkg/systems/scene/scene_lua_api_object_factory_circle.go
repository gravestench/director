package scene

import (
	lua "github.com/yuin/gopher-lua"

	"github.com/gravestench/director/pkg/util"
)

func (s *Scene) luaBindSceneObjectFactoryCircle(objFactoryTable *lua.LTable) {
	const name = "circle"

	fn := s.Lua.NewFunction(func(L *lua.LState) int {
		// check argument count
		if L.GetTop() != 5 {
			return 0
		}

		x := int(L.CheckNumber(1))
		y := int(L.CheckNumber(2))
		r := int(L.CheckNumber(3))

		fill, _ := util.ParseHexColor(L.CheckString(4))
		stroke, _ := util.ParseHexColor(L.CheckString(5))

		e := s.Add.Circle(x, y, r, fill, stroke)

		L.Push(lua.LNumber(e))

		return 1
	})

	s.Lua.SetField(objFactoryTable, name, fn)
}
