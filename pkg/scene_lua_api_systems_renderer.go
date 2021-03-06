package pkg

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	lua "github.com/yuin/gopher-lua"
)

const (
	luaRendererSystemName = "renderer"
)

func (s *SceneSystem) luaExportSystemRenderer(mt *lua.LTable) {
	sysTable := s.Lua.NewTypeMetatable(luaRendererSystemName)

	windowTable := s.Lua.NewTable()
	s.Lua.SetField(windowTable, "size", s.Lua.NewFunction(func(L *lua.LState) int {
		if s.Lua.GetTop() == 2 {
			w, h := int(s.Lua.CheckNumber(1)), int(s.Lua.CheckNumber(2))

			s.Director.Sys.Renderer.Window.Width, s.Director.Sys.Renderer.Window.Height = w, h
			return 0
		}

		s.Lua.Push(lua.LNumber(s.Director.Sys.Renderer.Window.Width))
		s.Lua.Push(lua.LNumber(s.Director.Sys.Renderer.Window.Height))

		return 2
	}))

	logTable := s.Lua.NewTable()
	makeLogLevelFunc := func(lvl int) *lua.LFunction {
		return s.Lua.NewFunction(func(L *lua.LState) int {
			s.Director.Sys.Renderer.Logging = lvl
			return 0
		})
	}

	s.Lua.SetField(logTable, "All", makeLogLevelFunc(rl.LogAll))
	s.Lua.SetField(logTable, "Trace", makeLogLevelFunc(rl.LogTrace))
	s.Lua.SetField(logTable, "Debug", makeLogLevelFunc(rl.LogDebug))
	s.Lua.SetField(logTable, "Info", makeLogLevelFunc(rl.LogInfo))
	s.Lua.SetField(logTable, "Warning", makeLogLevelFunc(rl.LogWarning))
	s.Lua.SetField(logTable, "Error", makeLogLevelFunc(rl.LogError))
	s.Lua.SetField(logTable, "Fatal", makeLogLevelFunc(rl.LogFatal))
	s.Lua.SetField(logTable, "None", makeLogLevelFunc(rl.LogNone))

	s.Lua.SetField(sysTable, "fps", s.Lua.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 1 {
			fpsSet := L.CheckNumber(1)
			s.Director.Sys.Renderer.TargetFPS = int(fpsSet)
		}

		L.Push(lua.LNumber(s.Director.Sys.Renderer.TargetFPS))

		return 1
	}))

	s.Lua.SetField(sysTable, "log", logTable)
	s.Lua.SetField(sysTable, "window", windowTable)

	s.Lua.SetField(mt, luaRendererSystemName, sysTable)
}
