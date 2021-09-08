package scene

import (
	"fmt"
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common"
	lua "github.com/yuin/gopher-lua"
)

const luaImageTypeName = "image"

var imageMethods = map[string]lua.LGFunction{
	"id": imageGet,
}

var luaImageTypeExporter = func(scene *Scene) common.LuaTypeExport {
	export := common.LuaTypeExport{
		Name:            luaImageTypeName,
		ConstructorFunc: scene.luaImageConstructor(),
		Methods:         imageMethods,
	}

	return export
}

func (s *Scene) luaImageConstructor() lua.LGFunction {
	return func(L *lua.LState) int {
		// check argument count
		if L.GetTop() != 3 {
			return 0
		}

		uri := L.CheckString(1)
		x := int(L.CheckNumber(2))
		y := int(L.CheckNumber(3))

		e := s.Add.Image(uri, x, y)
		v := &e

		ud := L.NewUserData()
		ud.Value = v
		L.SetMetatable(ud, L.GetTypeMetatable(luaImageTypeName))
		L.Push(ud)
		return 1
	}
}

// Checks whether the first Lua argument is a *LUserData with *Image and returns this *Image.
func checkImage(L *lua.LState) *akara.EID {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*akara.EID); ok {
		return v
	}
	L.ArgError(1, "image expected")
	return nil
}

// Getter and setter for the Image#XYZ
func imageGet(L *lua.LState) int {
	p := checkImage(L)

	L.Push(lua.LNumber(*p))

	return 1
}

func imageFromLua(ud *lua.LUserData) (*akara.EID, error) {
	if vv, ok := ud.Value.(*akara.EID); ok {
		return vv, nil
	}

	return nil, fmt.Errorf("failed to convert userdata to image")
}