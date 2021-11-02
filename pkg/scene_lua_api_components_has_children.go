package pkg

import (
	"strconv"

	lua "github.com/yuin/gopher-lua"

	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/components"
)

const (
	luaHasChildrenComponentName = "hasChildren"
)

/*
example lua:
	parent = newEntity()
	hc = components.hasChildren.add(parent)

	child1 = newEntity()
	child2 = newEntity()
	child3 = newEntity()

	hc.children({child1, child2, child3})
*/

func (s *Scene) luaExportComponentHasChildren(mt *lua.LTable) {
	hasChildrenTable := s.Lua.NewTypeMetatable(luaHasChildrenComponentName)

	s.Lua.SetField(hasChildrenTable, "add", s.Lua.NewFunction(s.luaHasChildrenAdd()))
	s.Lua.SetField(hasChildrenTable, "get", s.Lua.NewFunction(s.luaHasChildrenGet()))
	s.Lua.SetField(hasChildrenTable, "remove", s.Lua.NewFunction(s.luaHasChildrenRemove()))

	s.Lua.SetField(mt, luaHasChildrenComponentName, hasChildrenTable)
}

func (s *Scene) luaHasChildrenAdd() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		hasChildren := s.Components.HasChildren.Add(e)
		L.Push(s.makeLuaTableComponentHasChildren(hasChildren))
		return 1
	}

	return fn
}

func (s *Scene) luaHasChildrenGet() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		id := L.CheckNumber(1)
		hasChildren, found := s.Components.HasChildren.Get(common.Entity(id))

		truthy := lua.LFalse
		if !found {
			L.Push(lua.LNil)
			L.Push(truthy)
			return 2
		} else {
			truthy = lua.LTrue
		}

		table := s.makeLuaTableComponentHasChildren(hasChildren)

		L.SetMetatable(table, L.GetTypeMetatable(luaHasChildrenComponentName))

		L.Push(table)
		L.Push(truthy)

		return 2
	}

	return fn
}

func (s *Scene) luaHasChildrenRemove() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		s.Components.HasChildren.Remove(e)

		return 0
	}

	return fn
}

func (s *Scene) makeLuaTableComponentHasChildren(hc *components.HasChildren) *lua.LTable {
	table := s.Lua.NewTable()

	s.Lua.SetField(table, "children", s.Lua.NewFunction(func(L *lua.LState) int {
		numArgs := L.GetTop()

		// yield children in a table
		if numArgs < 1 {
			dataTable := L.NewTable()

			if hc.Children == nil {
				return 0
			}

			for idx, child := range hc.Children {
				L.SetField(dataTable, strconv.Itoa(idx), lua.LNumber(child))
			}

			L.Push(dataTable)

			return 1
		}

		// else, append children
		list := L.CheckTable(1)
		list.ForEach(func(idx, val lua.LValue) {
			eid, err := strconv.Atoi(val.String())
			if err != nil {
				return
			}

			hc.Children = append(hc.Children, common.Entity(eid))
		})

		return 0
	}))

	return table
}
