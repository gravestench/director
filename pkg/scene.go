package pkg

import (
	"github.com/gravestench/akara"
)

// SceneInterface represents what director considers to be a scene
type SceneInterface interface {
	akara.System
	HasKey
	initializesLua
	GenericUpdate()
	GenericSceneInit(d *Director)
	Render()
}

type HasKey interface {
	Key() string
}

type initializesLua interface {
	LuaInitialized() bool
	InitializeLua()
	UninitializeLua()
}
