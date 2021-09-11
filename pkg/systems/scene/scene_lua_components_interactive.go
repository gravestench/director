package scene

import (
	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/systems/input"
	lua "github.com/yuin/gopher-lua"
)

const (
	luaInteractiveComponentName = "interactive"
)

func (s *Scene) luaExportComponentInteractive(mt *lua.LTable) {
	interactiveTable := s.Lua.NewTypeMetatable(luaInteractiveComponentName)

	s.Lua.SetField(interactiveTable, "add", s.Lua.NewFunction(s.luaInteractiveAdd()))
	s.Lua.SetField(interactiveTable, "get", s.Lua.NewFunction(s.luaInteractiveGet()))

	s.Lua.SetField(mt, luaInteractiveComponentName, interactiveTable)
}

func (s *Scene) luaInteractiveAdd() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		interactive := s.Components.Interactive.Add(*s.luaCheckEID())
		L.Push(s.makeLuaTableComponentInteractive(interactive))
		return 1
	}

	return fn
}

func (s *Scene) luaInteractiveGet() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		id := L.CheckNumber(1)
		interactive, found := s.Components.Interactive.Get(common.Entity(id))

		truthy := lua.LFalse
		if !found {
			L.Push(lua.LNil)
			L.Push(truthy)
			return 2
		} else {
			truthy = lua.LTrue
		}

		table := s.makeLuaTableComponentInteractive(interactive)

		L.SetMetatable(table, L.GetTypeMetatable(luaInteractiveComponentName))

		L.Push(table)
		L.Push(truthy)

		return 2
	}

	return fn
}

func (s *Scene) makeLuaTableComponentInteractive(in *input.Interactive) *lua.LTable {
	table := s.Lua.NewTable()

	fnSetKey := func(L *lua.LState) int {
		if L.GetTop() < 1 {
			return 0
		}

		in.Vector.SetKey(input.Key(L.CheckNumber(1)))

		return 0
	}

	fnSetMod := func(L *lua.LState) int {
		if L.GetTop() < 1 {
			return 0
		}

		in.Vector.SetModifier(input.Modifier(L.CheckNumber(1)))

		return 0
	}

	fnSetMouseButton := func(L *lua.LState) int {
		if L.GetTop() < 1 {
			return 0
		}

		in.Vector.SetMouseButton(input.MouseButton(L.CheckNumber(1)))

		return 0
	}

	s.Lua.SetField(table, "setKey", s.Lua.NewFunction(fnSetKey))
	s.Lua.SetField(table, "setMod", s.Lua.NewFunction(fnSetMod))
	s.Lua.SetField(table, "setMouse", s.Lua.NewFunction(fnSetMouseButton))

	s.Lua.SetField(table, "hitbox", s.makeLuaTableImageRectangle(in.Hitbox))
	s.Lua.SetField(table, "callback", s.makeLuaInteractiveCallbackSetGet(in))

	return table
}
