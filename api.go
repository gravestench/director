package director

import (
	"flag"

	"github.com/gravestench/director/pkg"
	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/primitive_scenes/tick_graph"
)

// Director is the scene manager
type Director = pkg.Director

// Entity is the the base entity type
type Entity = common.Entity

// SceneComponents contains component factories that scenes use for composing objects
type SceneComponents = common.SceneComponents

// SceneInterface defines what a Scene is to director
type SceneInterface = pkg.SceneInterface

// SceneObjects contains object factories used by scenes for composing entities in the scene.
type SceneObjects = pkg.ObjectFactory

// System is an extension of an akara system, with director lifecycle methods,
// object factories, components, lua state machine, etc.
type System = pkg.Base

// Scene is an extension of a director system, which has director lifecycle methods,
// object factories, components, lua state machine, etc.
//
// The main difference between a Director System and a Scene is that a
// System is non-graphical, but still has access to most things a scene does.
type Scene = pkg.Scene

// PrimitiveSystems is a struct containing systems
type SceneSystems = pkg.DirectorSystems

// New returns a new director instance
func New() *Director {
	flag.Parse()

	instance := pkg.New()

	if flag.Lookup(pkg.FlagNameDebug).Value.(flag.Getter).Get().(bool) {
		instance.AddScene(&tick_graph.Scene{})
	}

	return instance
}

// NewLuaSystem creates a new system from a lua script. The lua script needs to have
// both `init` and `update(TimeDelta)` functions declared.
//
// A lua system is like a non-graphical scene, but still has access to some of the
// director scene stuff.
func NewLuaSystem(key, path string) *pkg.LuaSystem {
	return pkg.NewLuaSystem(key, path)
}

// NewLuaScene creates a new scene from a lua script. The lua script needs to have
// both `init` and `update(TimeDelta)` functions declared.
func NewLuaScene(key, path string) *pkg.LuaScene {
	return pkg.NewLuaScene(key, path)
}
