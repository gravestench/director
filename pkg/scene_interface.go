package pkg

import (
	"github.com/gravestench/akara"
)

// SceneInterface represents what director considers to be a scene
type SceneInterface interface {
	akara.System
	initializesLua
	isGeneric
	Key() string
	Render()
}

type isGeneric interface {
	GenericUpdate()
	GenericSceneInit(d *Director)
}

type initializesLua interface {
	LuaInitialized() bool
	InitializeLua()
	UninitializeLua()
}
