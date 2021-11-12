package pkg

import (
	lua "github.com/yuin/gopher-lua"

	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/components"
)

const (
	luaRenderOrderComponentName = "renderOrder"
)

/*
example lua:
	renderOrder = components.renderOrder.add(eid)

	-- the highest index is rendered on top of everything else
	index = renderOrder.value()
	renderOrder.value(100) -- set the layer index to 100
*/

func (s *SceneSystem) luaExportComponentRenderOrder(mt *lua.LTable) {
	renderOrderTable := s.Lua.NewTypeMetatable(luaRenderOrderComponentName)

	s.Lua.SetField(renderOrderTable, "add", s.Lua.NewFunction(s.luaRenderOrderAdd()))
	s.Lua.SetField(renderOrderTable, "get", s.Lua.NewFunction(s.luaRenderOrderGet()))
	s.Lua.SetField(renderOrderTable, "remove", s.Lua.NewFunction(s.luaRenderOrderRemove()))

	s.Lua.SetField(mt, luaRenderOrderComponentName, renderOrderTable)
}

func (s *SceneSystem) luaRenderOrderAdd() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		renderOrder := s.Components.RenderOrder.Add(e)
		L.Push(s.makeLuaTableComponentRenderOrder(renderOrder))
		return 1
	}

	return fn
}

func (s *SceneSystem) luaRenderOrderGet() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		id := L.CheckNumber(1)
		renderOrder, found := s.Components.RenderOrder.Get(common.Entity(id))

		truthy := lua.LFalse
		if !found {
			L.Push(lua.LNil)
			L.Push(truthy)
			return 2
		} else {
			truthy = lua.LTrue
		}

		table := s.makeLuaTableComponentRenderOrder(renderOrder)

		L.SetMetatable(table, L.GetTypeMetatable(luaRenderOrderComponentName))

		L.Push(table)
		L.Push(truthy)

		return 2
	}

	return fn
}

func (s *SceneSystem) luaRenderOrderRemove() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		s.Components.RenderOrder.Remove(e)

		return 0
	}

	return fn
}

func (s *SceneSystem) makeLuaTableComponentRenderOrder(renderOrder *components.RenderOrder) *lua.LTable {
	table := s.Lua.NewTable()

	s.Lua.SetField(table, "value", s.Lua.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 1 {
			val := int(L.CheckNumber(1))

			renderOrder.Value = val

			return 0
		}

		L.Push(lua.LNumber(renderOrder.Value))

		return 1
	}))

	return table
}
