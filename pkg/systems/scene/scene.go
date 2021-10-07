package scene

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	director "github.com/gravestench/director/pkg"
	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/mathlib"
	"github.com/gravestench/scenegraph"
	lua "github.com/yuin/gopher-lua"
	"math"
	"time"
)

const (
	luaSceneTable = "scene"
	luaSceneSystemsTable = "sys" // scene.sys
	luaSceneComponentsTable = "components" // scene.components
	luaSceneObjectFactoryTable = "add" // scene.add

	luaConstantsTable = "constants"
)

type Scene struct {
	akara.BaseSystem
	*director.Director
	Lua        *lua.LState
	Components common.BasicComponents
	Graph      scenegraph.Node
	key        string
	Add        ObjectFactory
	renderList []common.Entity
	Viewports  []common.Entity
}

var tmpVect mathlib.Vector3

func (s *Scene) renderEntity(e common.Entity) {
	texture, textureFound := s.Components.Texture2D.Get(e)
	rt, rtFound := s.Components.RenderTexture2D.Get(e)
	if !textureFound && !rtFound {
		return
	}

	var t *rl.Texture2D

	if !rtFound {
		t = texture.Texture2D
	} else {
		t = &rt.Texture
	}

	trs, found := s.Components.Transform.Get(e)
	if !found {
		return
	}

	origin, found := s.Components.Origin.Get(e)
	if !found {
		return
	}

	tint := rl.White

	opacity, found := s.Components.Opacity.Get(e)
	if found {
		if opacity.Value > 1 {
			opacity.Value = 1
		} else if opacity.Value < 0 {
			opacity.Value = 0
		}

		tint.A = uint8(float64(math.MaxUint8) * opacity.Value)
	}

	if tint.A == 0 {
		return
	}

	// this is rotating around the origin point from the origin component
	tmpVect.Set(float64(t.Width), float64(t.Height), 1)
	yRad := trs.Rotation.Y * mathlib.DegreesToRadians
	ov2 := mathlib.NewVector2(origin.Clone().Multiply(&tmpVect).XY()).Rotate(yRad).Negate()
	ov3 := mathlib.NewVector3(ov2.X, ov2.Y, 0)
	x, y := trs.Translation.Clone().Add(ov3).XY()
	v2 := mathlib.NewVector2(x, y)

	position := rl.Vector2{X: float32(v2.X), Y: float32(v2.Y)}
	rotation := float32(trs.Rotation.Y)
	scale := float32(trs.Scale.X)

	rl.DrawTextureEx(*t, position, rotation, scale, tint)
}

func (s *Scene) GenericSceneInit(d *director.Director) {
	s.Add.scene = s
	s.Director = d
	s.Components.Init(s.Director.World)
	s.BaseSystem.SetPreTickFunc(func() {
		s.GenericUpdate()
	})
}

func (s *Scene) InitializeLua() {
	if s.LuaInitialized() {
		return
	}

	s.Lua = lua.NewState()
	if err := s.Lua.DoString(common.LuaLibSTD); err != nil {
		panic(err)
	}

	s.initLuaConstantsTable()
	s.initLuaSceneTable()
}

func (s *Scene) UninitializeLua() {
	s.Lua = nil
}

func (s *Scene) LuaInitialized() bool {
	return s.Lua != nil
}

func (s *Scene) updateSceneGraph() {
	for _, eid := range s.renderList {
		node, found := s.Components.SceneGraphNode.Get(eid)
		if !found {
			continue
		}

		trs, found := s.Components.Transform.Get(eid)
		if !found {
			continue
		}

		node.Local = trs.GetMatrix()
	}

	s.Graph.UpdateWorldMatrix()
}

func (s *Scene) updateSceneObjects(dt time.Duration) {
	s.Add.update(dt)
}

func (s *Scene) GenericUpdate() {
	s.updateSceneGraph()
	s.updateSceneObjects(s.Director.TimeDelta)

	// this renders the scene objects to the scene's render texture
	// however, this will not actually display anything, that is done by the render system
	s.Render()
}

func (s *Scene) Render() {
	if len(s.Viewports) < 1 {
		s.initViewport()
	}

	for _, e := range s.Viewports {
		s.renderCameraForViewport(e)
		s.renderCameraToViewport(e)
	}
}

func (s *Scene) RemoveEntity(e common.Entity) {
	factories := []*akara.ComponentFactory{
		s.Components.Viewport.ComponentFactory,
		s.Components.Camera.ComponentFactory,
		s.Components.Color.ComponentFactory,
		s.Components.Debug.ComponentFactory,
		s.Components.FileLoadRequest.ComponentFactory,
		s.Components.FileLoadResponse.ComponentFactory,
		s.Components.FileType.ComponentFactory,
		s.Components.Fill.ComponentFactory,
		s.Components.Animation.ComponentFactory,
		s.Components.Origin.ComponentFactory,
		s.Components.Opacity.ComponentFactory,
		s.Components.Stroke.ComponentFactory,
		s.Components.Font.ComponentFactory,
		s.Components.SceneGraphNode.ComponentFactory,
		s.Components.Text.ComponentFactory,
		s.Components.RenderTexture2D.ComponentFactory,
		s.Components.Size.ComponentFactory,
		s.Components.Texture2D.ComponentFactory,
		s.Components.Transform.ComponentFactory,
		s.Components.UUID.ComponentFactory,
	}

	// nuke all components the entity may have
	for idx := range factories {
		factories[idx].Remove(e)
	}


	for idx := range s.renderList {
		if s.renderList[idx] == e {
			s.renderList = append(s.renderList[:idx], s.renderList[idx+1:]...)
			break
		}
	}


	s.Director.RemoveEntity(e)
}
