package scene

import (
	"github.com/gravestench/director/pkg/util"

	"github.com/gravestench/director/pkg/common"
	lua "github.com/yuin/gopher-lua"
)

const luaRectangleTypeName = "rectangle"

var rectangleMethods = map[string]lua.LGFunction{
	"id": rectangleGet,
}

var luaRectangleTypeExporter = func(scene *Scene) common.LuaTypeExport {
	export := common.LuaTypeExport{
		Name:            luaRectangleTypeName,
		ConstructorFunc: scene.luaRectangleConstructor(),
		Methods:         rectangleMethods,
	}

	return export
}

func (s *Scene) luaRectangleConstructor() lua.LGFunction {
	return func(L *lua.LState) int {
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
		v := &e

		ud := L.NewUserData()
		ud.Value = v
		L.SetMetatable(ud, L.GetTypeMetatable(luaRectangleTypeName))
		L.Push(ud)
		return 1
	}
}

// Checks whether the first Lua argument is a *LUserData with *Rectangle and returns this *Rectangle.
func checkRectangle(L *lua.LState) *common.Entity {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*common.Entity); ok {
		return v
	}
	L.ArgError(1, "rectangle expected")
	return nil
}

// Getter and setter for the Rectangle#XYZ
func rectangleGet(L *lua.LState) int {
	p := checkRectangle(L)

	L.Push(lua.LNumber(*p))

	return 1
}