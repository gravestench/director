package scene

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	lua "github.com/yuin/gopher-lua"
)

const (
	luaConstantsLogName = "log"
)

func (s *Scene) luaExportConstantsLogging(mt *lua.LTable) {
	keyTable := s.Lua.NewTypeMetatable(luaConstantsLogName)

	keys := map[string]int{
		"all":     rl.LogAll,
		"trace":   rl.LogTrace,
		"debug":   rl.LogDebug,
		"info":    rl.LogInfo,
		"warning": rl.LogWarning,
		"error":   rl.LogError,
		"fatal":   rl.LogFatal,
		"none":    rl.LogNone,
	}

	for k, v := range keys {
		s.Lua.SetField(keyTable, k, lua.LNumber(v))
	}

	s.Lua.SetField(mt, luaConstantsLogName, keyTable)
}
