package scene

import (
	"github.com/gravestench/director/pkg/util"
	lua "github.com/yuin/gopher-lua"
	"image/color"
)

func (s *Scene) luaBindSceneObjectFactoryLabel(objFactoryTable *lua.LTable) {
	const name = "label"

	fn := s.Lua.NewFunction(func(L *lua.LState) int {
		// check argument count, this is coming from inside of lua
		if L.GetTop() != 6 {
			return 0
		}

		// pop the values that were passed into the function
		text := L.CheckString(1)
		x := int(L.CheckNumber(2))
		y := int(L.CheckNumber(3))
		size := int(L.CheckNumber(4))
		fontName := L.CheckString(5)
		c, err := util.ParseHexColor(L.CheckString(6))
		if err != nil {
			c = color.RGBA{R: 255, B: 255, A: 196}
		}

		// use them with the object factory
		e := s.Add.Label(text, x, y, size, fontName, c)

		L.Push(lua.LNumber(e))

		return 1
	})

	s.Lua.SetField(objFactoryTable, name, fn)
}
