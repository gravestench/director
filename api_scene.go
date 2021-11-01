package director

import (
	"github.com/gravestench/director/pkg"
	"github.com/gravestench/director/pkg/systems/scene"
	"github.com/gravestench/director/primitive_scenes/tick_graph"
)

type SceneInterface = pkg.SceneInterface

// Scene, a director API namespace, that contains anything an end-user should need from ~/pkg regarding scene systems.
//
// This includes provider functions, as well as other utility functions, that
// any end-user of Director would want to know about without having to go dig around
// inside of pkg.
//
// The main rationale is that this is just a nice way of exporting scene-related stuff
// from inside of pkg. The things that get exported might not all be from inside
// of ~/pkg/systems/scene/
var Scene sceneNamespace

type sceneNamespace struct {} // a stateless singleton

func (sceneNamespace) NewLuaScene(sceneKey, luaFilePath string) *scene.LuaScene {
	return scene.NewLuaScene(sceneKey, luaFilePath)
}

const (
	ScenePrimitiveTickGraph = tick_graph.Key
)

// Primitives returns primitive scenes which are included with director.
// These are general-purpose scenes that can be used by end-users, but are not added to director by default.
//
// These can be for debug, interactive terminals, etc.
func (sceneNamespace) Primitives() map[string]pkg.SceneInterface {
	primitives := map[string]pkg.SceneInterface{
		ScenePrimitiveTickGraph: &tick_graph.Scene{}, // for plotting scene tick rate in a simple graph on screen
	}

	return primitives
}