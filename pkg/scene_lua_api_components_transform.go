package pkg

import (
	lua "github.com/yuin/gopher-lua"

	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/components"
)

const (
	luaTransformComponentName = "transform"
)

func (s *SceneSystem) luaExportComponentTransform(mt *lua.LTable) {
	trsTable := s.Lua.NewTypeMetatable(luaTransformComponentName)

	s.Lua.SetField(trsTable, "add", s.Lua.NewFunction(s.luaTransformAdd()))
	s.Lua.SetField(trsTable, "get", s.Lua.NewFunction(s.luaTransformGet()))
	s.Lua.SetField(trsTable, "remove", s.Lua.NewFunction(s.luaTransformRemove()))

	s.Lua.SetField(mt, luaTransformComponentName, trsTable)
}

func (s *SceneSystem) luaTransformAdd() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		trs := s.Components.Transform.Add(e)
		L.Push(s.makeLuaTableComponentTransform(trs))
		return 1
	}

	return fn
}

func (s *SceneSystem) luaTransformGet() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		id := L.CheckNumber(1)
		trs, found := s.Components.Transform.Get(common.Entity(id))

		truthy := lua.LFalse
		if !found {
			L.Push(lua.LNil)
			L.Push(truthy)
			return 2
		} else {
			truthy = lua.LTrue
		}

		table := s.makeLuaTableComponentTransform(trs)

		L.SetMetatable(table, L.GetTypeMetatable(luaTransformComponentName))

		L.Push(table)
		L.Push(truthy)

		return 2
	}

	return fn
}

func (s *SceneSystem) luaTransformRemove() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		s.Components.Transform.Remove(e)

		return 0
	}

	return fn
}

func (s *SceneSystem) makeLuaTableComponentTransform(trs *components.Transform) *lua.LTable {
	table := s.Lua.NewTable()

	s.Lua.SetField(table, "translation", s.makeLuaSetterGetterVec3(trs.Translation))
	s.Lua.SetField(table, "rotation", s.makeLuaSetterGetterVec3(trs.Rotation))
	s.Lua.SetField(table, "scale", s.makeLuaSetterGetterVec3(trs.Scale))

	return table
}
