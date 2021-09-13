package scene

import (
	lua "github.com/yuin/gopher-lua"
	"image/color"

	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/components"
)

const (
	luaStrokeComponentName = "stroke"
)

/*
example lua:
	stroke = scene.components.stroke.add(eid)

	stroke.rgba(255, 255, 255, 255)
	r, g, b, a = stroke.rgba()
*/

func (s *Scene) luaExportComponentStroke(mt *lua.LTable) {
	strokeTable := s.Lua.NewTypeMetatable(luaStrokeComponentName)

	s.Lua.SetField(strokeTable, "add", s.Lua.NewFunction(s.luaStrokeAdd()))
	s.Lua.SetField(strokeTable, "get", s.Lua.NewFunction(s.luaStrokeGet()))
	s.Lua.SetField(strokeTable, "remove", s.Lua.NewFunction(s.luaStrokeRemove()))

	s.Lua.SetField(mt, luaStrokeComponentName, strokeTable)
}

func (s *Scene) luaStrokeAdd() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		c := s.Components.Stroke.Add(e)
		L.Push(s.makeLuaTableComponentStroke(c))
		return 1
	}

	return fn
}

func (s *Scene) luaStrokeGet() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		id := L.CheckNumber(1)
		stroke, found := s.Components.Stroke.Get(common.Entity(id))

		truthy := lua.LFalse
		if !found {
			L.Push(lua.LNil)
			L.Push(truthy)
			return 2
		} else {
			truthy = lua.LTrue
		}

		table := s.makeLuaTableComponentStroke(stroke)

		L.SetMetatable(table, L.GetTypeMetatable(luaStrokeComponentName))

		L.Push(table)
		L.Push(truthy)

		return 2
	}

	return fn
}

func (s *Scene) luaStrokeRemove() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		s.Components.Stroke.Remove(e)

		return 0
	}

	return fn
}

func (s *Scene) makeLuaTableComponentStroke(stroke *components.Stroke) *lua.LTable {
	table := s.Lua.NewTable()

	s.Lua.SetField(table, "rgba", s.Lua.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 4 {
			r := uint8(L.CheckNumber(1))
			g := uint8(L.CheckNumber(1))
			b := uint8(L.CheckNumber(1))
			a := uint8(L.CheckNumber(1))

			stroke.Color = color.RGBA{
				R: r,
				G: g,
				B: b,
				A: a,
			}

			return 0
		}

		r, g, b, a := stroke.RGBA()

		L.Push(lua.LNumber(r))
		L.Push(lua.LNumber(g))
		L.Push(lua.LNumber(b))
		L.Push(lua.LNumber(a))

		return 4
	}))

	return table
}
