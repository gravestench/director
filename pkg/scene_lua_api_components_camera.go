package pkg

import (
	lua "github.com/yuin/gopher-lua"

	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/components"
)

const (
	luaCameraComponentName = "camera"
)

/*
example lua:
	cam = components.camera.add(eid)

	cam.zoom(1.5)
	zoom = cam.zoom()

	cam.rotate(90)
	rot = cam.rotate()
*/

func (s *Scene) luaExportComponentCamera(mt *lua.LTable) {
	cameraTable := s.Lua.NewTypeMetatable(luaCameraComponentName)

	s.Lua.SetField(cameraTable, "add", s.Lua.NewFunction(s.luaCameraAdd()))
	s.Lua.SetField(cameraTable, "get", s.Lua.NewFunction(s.luaCameraGet()))
	s.Lua.SetField(cameraTable, "remove", s.Lua.NewFunction(s.luaCameraRemove()))

	s.Lua.SetField(mt, luaCameraComponentName, cameraTable)
}

func (s *Scene) luaCameraAdd() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		camera := s.Components.Camera.Add(e)
		L.Push(s.makeLuaTableComponentCamera(camera))
		return 1
	}

	return fn
}

func (s *Scene) luaCameraGet() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		id := L.CheckNumber(1)
		camera, found := s.Components.Camera.Get(common.Entity(id))

		truthy := lua.LFalse
		if !found {
			L.Push(lua.LNil)
			L.Push(truthy)
			return 2
		} else {
			truthy = lua.LTrue
		}

		table := s.makeLuaTableComponentCamera(camera)

		L.SetMetatable(table, L.GetTypeMetatable(luaCameraComponentName))

		L.Push(table)
		L.Push(truthy)

		return 2
	}

	return fn
}

func (s *Scene) luaCameraRemove() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		s.Components.Camera.Remove(e)

		return 0
	}

	return fn
}

func (s *Scene) makeLuaTableComponentCamera(camera *components.Camera) *lua.LTable {
	table := s.Lua.NewTable()

	s.Lua.SetField(table, "rotation", s.Lua.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 0 {
			L.Push(lua.LNumber(camera.Rotation))

			return 1
		}

		r := float32(L.CheckNumber(1))
		camera.Rotation = r

		return 0
	}))

	s.Lua.SetField(table, "zoom", s.Lua.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 0 {
			L.Push(lua.LNumber(camera.Zoom))

			return 1
		}

		z := float32(L.CheckNumber(1))
		camera.Zoom = z

		return 0
	}))

	return table
}
