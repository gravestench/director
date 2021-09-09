package scene

import (
	lua "github.com/yuin/gopher-lua"

	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/util"
)

const luaCircleTypeName = "circle"

var circleMethods = map[string]lua.LGFunction{
	"id": circleGet,
}

var luaCircleTypeExporter = func(scene *Scene) common.LuaTypeExport {
	export := common.LuaTypeExport{
		Name:            luaCircleTypeName,
		ConstructorFunc: scene.luaCircleConstructor(),
		Methods:         circleMethods,
	}

	return export
}

func (s *Scene) luaCircleConstructor() lua.LGFunction {
	return func(L *lua.LState) int {
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
		v := &e

		ud := L.NewUserData()
		ud.Value = v
		L.SetMetatable(ud, L.GetTypeMetatable(luaCircleTypeName))
		L.Push(ud)
		return 1
	}
}

// Checks whether the first Lua argument is a *LUserData with *Circle and returns this *Circle.
func checkCircle(L *lua.LState) *common.Entity {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*common.Entity); ok {
		return v
	}
	L.ArgError(1, "circle expected")
	return nil
}

// Getter and setter for the Circle#XYZ
func circleGet(L *lua.LState) int {
	p := checkCircle(L)

	L.Push(lua.LNumber(*p))

	return 1
}