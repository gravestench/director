package scene

import (
	lua "github.com/yuin/gopher-lua"
	"io"
	"strconv"

	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/components"
)

const (
	luaFileLoadResponseComponentName = "fileLoadResponse"
)

/*
example lua:
	res = scene.components.fileLoadResponse.add(eid)

	bytes = res.data() -- raw byte stream
    for idx, byte in ipairs(bytes) do
        -- do something with the indexes and bytes
    end
*/

func (s *Scene) luaExportComponentFileLoadResponse(mt *lua.LTable) {
	fileLoadResponseTable := s.Lua.NewTypeMetatable(luaFileLoadResponseComponentName)

	s.Lua.SetField(fileLoadResponseTable, "add", s.Lua.NewFunction(s.luaFileLoadResponseAdd()))
	s.Lua.SetField(fileLoadResponseTable, "get", s.Lua.NewFunction(s.luaFileLoadResponseGet()))
	s.Lua.SetField(fileLoadResponseTable, "remove", s.Lua.NewFunction(s.luaFileLoadResponseRemove()))

	s.Lua.SetField(mt, luaFileLoadResponseComponentName, fileLoadResponseTable)
}

func (s *Scene) luaFileLoadResponseAdd() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		fileLoadResponse := s.Components.FileLoadResponse.Add(e)
		L.Push(s.makeLuaTableComponentFileLoadResponse(fileLoadResponse))
		return 1
	}

	return fn
}

func (s *Scene) luaFileLoadResponseGet() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		id := L.CheckNumber(1)
		fileLoadResponse, found := s.Components.FileLoadResponse.Get(common.Entity(id))

		truthy := lua.LFalse
		if !found {
			L.Push(lua.LNil)
			L.Push(truthy)
			return 2
		} else {
			truthy = lua.LTrue
		}

		table := s.makeLuaTableComponentFileLoadResponse(fileLoadResponse)

		L.SetMetatable(table, L.GetTypeMetatable(luaFileLoadResponseComponentName))

		L.Push(table)
		L.Push(truthy)

		return 2
	}

	return fn
}

func (s *Scene) luaFileLoadResponseRemove() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		s.Components.FileLoadResponse.Remove(e)

		return 0
	}

	return fn
}

func (s *Scene) makeLuaTableComponentFileLoadResponse(res *components.FileLoadResponse) *lua.LTable {
	table := s.Lua.NewTable()

	s.Lua.SetField(table, "data", s.Lua.NewFunction(func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		dataTable := L.NewTable()

		if res.Stream == nil {
			return 0
		}

		_, err := res.Stream.Seek(0, io.SeekStart)
		if err != nil {
			return 0
		}

		bytes, err := io.ReadAll(res.Stream)
		if err != nil {
			return 0
		}

		for idx := range bytes {
			L.SetField(dataTable, strconv.Itoa(idx), lua.LNumber(bytes[idx]))
		}

		L.Push(dataTable)

		return 1
	}))

	return table
}
