package vector3

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/director/cmd/examples/lua_api/common"
	lua "github.com/yuin/gopher-lua"
)

const luaVector3TypeName = "vector3"

var LuaTypeExport = common.LuaTypeExport{
	Name:            luaVector3TypeName,
	ConstructorFunc: newVector3,
	Methods: map[string]lua.LGFunction{
		"xyz":  vector3GetSetXYZ,  // from Lua: dog:name()
	},
}

// Constructor
func newVector3(L *lua.LState) int {
	if L.GetTop() != 3 {
		return 0
	}

	v := &rl.Vector3{
		X: float32(L.CheckNumber(1)),
		Y: float32(L.CheckNumber(2)),
		Z: float32(L.CheckNumber(3)),
	}

	ud := L.NewUserData()
	ud.Value = v
	L.SetMetatable(ud, L.GetTypeMetatable(luaVector3TypeName))
	L.Push(ud)
	return 1
}

// Checks whether the first lua argument is a *LUserData with *Vector3 and returns this *Vector3.
func checkVector3(L *lua.LState) *rl.Vector3 {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*rl.Vector3); ok {
		return v
	}
	L.ArgError(1, "vector3 expected")
	return nil
}

var vector3Methods = map[string]lua.LGFunction{
	"name": vector3GetSetXYZ,
}

// Getter and setter for the Vector3#XYZ
func vector3GetSetXYZ(L *lua.LState) int {
	p := checkVector3(L)

	if L.GetTop() == 4 {
		p.X = float32(L.CheckNumber(2))
		p.Y = float32(L.CheckNumber(3))
		p.Z = float32(L.CheckNumber(4))
		return 0
	}

	L.Push(lua.LNumber(p.X))
	L.Push(lua.LNumber(p.Y))
	L.Push(lua.LNumber(p.Z))

	return 3
}

func FromLua(ud *lua.LUserData) (*rl.Vector3, error) {
	if vv, ok := ud.Value.(*rl.Vector3); ok {
		return vv, nil
	}

	return nil, fmt.Errorf("failed to convert userdata to vector3")
}