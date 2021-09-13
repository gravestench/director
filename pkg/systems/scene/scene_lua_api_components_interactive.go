package scene

import (
	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/systems/input"
	lua "github.com/yuin/gopher-lua"
)

/*
example lua:
	i = scene.components.interactive.add(eid)

	-- set up a callack on mouse click
	-- in rectangle at (10,10), with dimensions 40x40
	i.setMouse(constants.input.MouseButtonLeft)
	i.hitbox(10, 10, 40, 40)
	i.callback("testCallback") -- notice that it is a string

	function testCallback()
		print("hello from lua")
	end
*/

func (s *Scene) luaExportComponentInteractive(mt *lua.LTable) {
	const name = "interactive"
	
	cTable := s.Lua.NewTable()

	s.Lua.SetField(cTable, "add", s.Lua.NewFunction(func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		interactive := s.Components.Interactive.Add(e)
		L.Push(s.makeLuaTableComponentInteractive(interactive))
		return 1
	}))

	s.Lua.SetField(cTable, "remove", s.Lua.NewFunction(func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		s.Components.Interactive.Remove(e)

		return 1
	}))

	s.Lua.SetField(cTable, "get", s.Lua.NewFunction(func(L *lua.LState) int {
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

		L.SetMetatable(table, L.GetTypeMetatable(name))

		L.Push(table)
		L.Push(truthy)

		return 2
	}))

	s.Lua.SetField(mt, name, cTable)
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
