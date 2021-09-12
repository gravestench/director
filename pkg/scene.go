package pkg

import (
	"github.com/gravestench/akara"
	"time"
)

// SceneInterface represents what director considers to be a scene
type SceneInterface interface {
	akara.System
	HasKey
	initializesLua
	GenericUpdate(duration time.Duration)
	GenericSceneInit(d *Director, width int, height int)
	Render()
}

type Updater interface {
	Update()
}

type UpdaterTimed interface {
	Update(duration time.Duration)
}

type HasKey interface {
	Key() string
}

type initializesLua interface {
	LuaInitialized() bool
	InitializeLua()
	UninitializeLua()
}
