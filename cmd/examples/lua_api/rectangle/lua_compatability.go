package rectangle

import (
	"fmt"
	"image"

	lua "github.com/yuin/gopher-lua"

	"github.com/gravestench/director/cmd/examples/lua_api/common"
)

// these two variables set up our Go struct for use in Lua scripts
var luaTypeExportName = "rectangle" // this sets the global name used to access this type from Lua

var LuaTypeExport = common.LuaTypeExport{
	Name:            luaTypeExportName,
	ConstructorFunc: newRectangleLua,
	Methods: map[string]lua.LGFunction{
		"name":  nameGetterSetter,     // from Lua: rectangle:name()
		"speak": positionGetterSetter, // from Lua: rectangle:speak()
	},
}

// newRectangleLua is the constructor for Rectangle objects in Lua scripts.
// Usage: rectangle.new(name)
// Example: my_square = rectangle.new("Rufus")
func newRectangleLua(L *lua.LState) int {
	x, y, w, h := L.CheckInt(1), L.CheckInt(2), L.CheckInt(3), L.CheckInt(4)

	square := &Rectangle{
		Rectangle: image.Rectangle{
			Min: image.Point{X: x, Y: y},
			Max: image.Point{X: x+w, Y: y+h},
		},
	}

	ud := L.NewUserData()
	ud.Value = square

	L.SetMetatable(ud, L.GetTypeMetatable(luaTypeExportName))
	L.Push(ud)

	return 1
}

func FromLua(ud *lua.LUserData) (*Rectangle, error) {
	if vv, ok := ud.Value.(*Rectangle); ok {
		return vv, nil
	}

	return nil, fmt.Errorf("failed to convert userdata to Rectangle")
}

// checkRectangle is a Go utility function that checks whether the first lua argument is a *LUserData representing a *Rectangle and
// returns this *Rectangle.
// This allows us to reliably translate a Lua Rectangle to a Go Rectangle.
func checkRectangle(L *lua.LState) *Rectangle {
	ud := L.CheckUserData(1)

	square, err := FromLua(ud)
	if err != nil {
		L.ArgError(1, "rectangle expected")
		return nil
	}

	return square
}

// nameGetterSetter is a combined getter and setter for Rectangle's Name field.
// Examples from Lua:
// Getter: my_square:name()
// Setter: my_square:name("Barney")
func nameGetterSetter(L *lua.LState) int {
	square := checkRectangle(L)

	// setter
	if L.GetTop() == 2 {
		square.Name = L.CheckString(2)

		return 0
	}

	// getter
	L.Push(lua.LString(square.Name))

	return 1
}

// positionGetterSetter is a combined getter and setter for Rectangle's x,y coordinate.
// Examples from Lua:
// Getter: 	my_square:speak()
// Setter: 	my_square:speak(function ()
// 				print("Bark bark!")
// 			end)
func positionGetterSetter(L *lua.LState) int {
	square := checkRectangle(L)

	// setter
	if L.GetTop() == 4 {
		// wrap the Lua function in some Go glue that allows it to run properly
		luaFunc := L.CheckFunction(2)
		square.Speak = func() {
			if err := L.CallByParam(lua.P{
				Fn: luaFunc,
				NRet: 1,
				Protect: true,
			}); err != nil {
				panic(err)
			}
		}

		return 0
	}

	// getter
	L.Push(L.NewFunction(func(L *lua.LState) int {
		square.Speak()
		return 0
	}))

	return 1
}
