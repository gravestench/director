package pkg

import (
	"github.com/gravestench/director/pkg/systems/audio"
	"github.com/gravestench/director/pkg/systems/file_loader"
	"github.com/gravestench/director/pkg/systems/input"
	"github.com/gravestench/director/pkg/systems/renderer"
	"github.com/gravestench/director/pkg/systems/texture_manager"
	"github.com/gravestench/director/pkg/systems/tween"
	lua "github.com/yuin/gopher-lua"
)

// DirectorSystems contains the base systems that are available when a director instance is created
type DirectorSystems struct {
	Load     *file_loader.System
	Renderer *renderer.System
	Texture  *texture_manager.System
	Tweens   *tween.System
	Input    *input.System
	Audio    *audio.System
}

func (d DirectorSystems) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	panic("implement me")
}
