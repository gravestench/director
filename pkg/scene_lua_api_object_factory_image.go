package pkg

import (
	lua "github.com/yuin/gopher-lua"
)

func (s *SceneSystem) luaBindSceneObjectFactoryImage(objFactoryTable *lua.LTable) {
	const name = "image"

	fn := s.Lua.NewFunction(func(L *lua.LState) int {
		// check argument count, this is coming from inside of lua
		if L.GetTop() != 3 {
			return 0
		}

		// pop the values that were passed into the function
		uri := L.CheckString(1)
		x := int(L.CheckNumber(2))
		y := int(L.CheckNumber(3))

		// use them with the object factory
		e := s.Add.Image(uri, x, y)

		L.Push(lua.LNumber(e))

		return 1
	})

	s.Lua.SetField(objFactoryTable, name, fn)
}
