package scene

import (
	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/util"
	lua "github.com/yuin/gopher-lua"
	"image/color"
)

const luaLabelTypeName = "label"

var labelMethods = map[string]lua.LGFunction{
	"id": labelGet,
}

var luaLabelTypeExporter = func(scene *Scene) common.LuaTypeExport {
	export := common.LuaTypeExport{
		Name:            luaLabelTypeName,
		ConstructorFunc: scene.luaLabelConstructor(),
		Methods:         labelMethods,
	}

	return export
}

func (s *Scene) luaLabelConstructor() lua.LGFunction {
	return func(L *lua.LState) int {
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
			c = color.RGBA{R: 255, B:255, A: 196}
		}

		// use them with the object factory
		e := s.Add.Label(text, x, y, size, fontName, c)
		v := &e

		ud := L.NewUserData()
		ud.Value = v
		L.SetMetatable(ud, L.GetTypeMetatable(luaLabelTypeName))
		L.Push(ud)
		return 1
	}
}

// Checks whether the first Lua argument is a *LUserData with *Label and returns this *Label.
func checkLabel(L *lua.LState) *common.Entity {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*common.Entity); ok {
		return v
	}
	L.ArgError(1, "label expected")
	return nil
}

// Getter and setter for the Label#XYZ
func labelGet(L *lua.LState) int {
	p := checkLabel(L)

	L.Push(lua.LNumber(*p))

	return 1
}
