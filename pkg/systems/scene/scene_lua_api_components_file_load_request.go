package scene

import (
	lua "github.com/yuin/gopher-lua"

	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/components"
)

const (
	luaFileLoadRequestComponentName = "fileLoadRequest"
)

/*
example lua:
	req = scene.components.fileLoadRequest.add(eid)

	path = req.Path()
	req.path("http://example.com/image.jpg")
*/

func (s *Scene) luaExportComponentFileLoadRequest(mt *lua.LTable) {
	fileLoadRequestTable := s.Lua.NewTypeMetatable(luaFileLoadRequestComponentName)

	s.Lua.SetField(fileLoadRequestTable, "add", s.Lua.NewFunction(s.luaFileLoadRequestAdd()))
	s.Lua.SetField(fileLoadRequestTable, "get", s.Lua.NewFunction(s.luaFileLoadRequestGet()))
	s.Lua.SetField(fileLoadRequestTable, "remove", s.Lua.NewFunction(s.luaFileLoadRequestRemove()))

	s.Lua.SetField(mt, luaFileLoadRequestComponentName, fileLoadRequestTable)
}

func (s *Scene) luaFileLoadRequestAdd() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		fileLoadRequest := s.Components.FileLoadRequest.Add(e)
		L.Push(s.makeLuaTableComponentFileLoadRequest(fileLoadRequest))
		return 1
	}

	return fn
}

func (s *Scene) luaFileLoadRequestGet() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		id := L.CheckNumber(1)
		fileLoadRequest, found := s.Components.FileLoadRequest.Get(common.Entity(id))

		truthy := lua.LFalse
		if !found {
			L.Push(lua.LNil)
			L.Push(truthy)
			return 2
		} else {
			truthy = lua.LTrue
		}

		table := s.makeLuaTableComponentFileLoadRequest(fileLoadRequest)

		L.SetMetatable(table, L.GetTypeMetatable(luaFileLoadRequestComponentName))

		L.Push(table)
		L.Push(truthy)

		return 2
	}

	return fn
}

func (s *Scene) luaFileLoadRequestRemove() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		s.Components.FileLoadRequest.Remove(e)

		return 0
	}

	return fn
}

func (s *Scene) makeLuaTableComponentFileLoadRequest(req *components.FileLoadRequest) *lua.LTable {
	table := s.Lua.NewTable()

	s.Lua.SetField(table, "path", s.Lua.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 1 {
			uri := L.CheckString(1)

			req.Path = uri

			return 0
		}

		L.Push(lua.LString(req.Path))

		return 1
	}))

	return table
}
