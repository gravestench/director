package pkg

import (
	lua "github.com/yuin/gopher-lua"

	"github.com/gravestench/director/pkg/util"
)

func (s *SceneSystem) luaBindSceneObjectFactoryRectangle(objFactoryTable *lua.LTable) {
	const name = "rectangle"

	fn := s.Lua.NewFunction(func(L *lua.LState) int {
		// check argument count
		if L.GetTop() != 6 {
			return 0
		}

		x := int(L.CheckNumber(1))
		y := int(L.CheckNumber(2))
		w := int(L.CheckNumber(3))
		h := int(L.CheckNumber(4))

		fill, _ := util.ParseHexColor(L.CheckString(5))
		stroke, _ := util.ParseHexColor(L.CheckString(6))

		e := s.Add.Rectangle(x, y, w, h, fill, stroke)

		L.Push(lua.LNumber(e))

		return 1
	})

	s.Lua.SetField(objFactoryTable, name, fn)
}
