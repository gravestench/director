package scene

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/mathlib"
	"github.com/gravestench/scenegraph"
	lua "github.com/yuin/gopher-lua"
	"time"
)

type Scene struct {
	akara.BaseSystem
	Lua           *lua.LState
	Components    common.BasicComponents
	Graph         scenegraph.Node
	key           string
	Add           ObjectFactory
	Renderables   *akara.Subscription
	Cameras       []akara.EID
	Width         int
	Height        int
}

var tmpVect mathlib.Vector3

func (s *Scene) renderEntity(e akara.EID) {
	texture, textureFound := s.Components.Texture2D.Get(e)
	rt, rtFound := s.Components.RenderTexture2D.Get(e)
	if !textureFound && !rtFound {
		return
	}

	var t *rl.Texture2D

	if !rtFound {
		t = &texture.Texture2D
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

	tmpVect.Set(float64(rt.Texture.Width), float64(rt.Texture.Height), 1)

	yRad := trs.Rotation.Y * mathlib.DegreesToRadians
	ov2 := mathlib.NewVector2(origin.Clone().Multiply(&tmpVect).XY()).Rotate(yRad).Negate()
	ov3 := mathlib.NewVector3(ov2.X, ov2.Y, 0)

	x, y := trs.Translation.Clone().Add(ov3).XY()
	v2 := mathlib.NewVector2(x, y)

	position := rl.Vector2{
		X: float32(v2.X),
		Y: float32(v2.Y),
	}

	rotation := float32(trs.Rotation.Y)

	scale := float32(trs.Scale.X)

	rl.DrawTextureEx(*t, position, rotation, scale, rl.White)
}

func (s *Scene) Initialize(width, height int, world *akara.World, renderablesSubscription *akara.Subscription) {
	s.Add.scene = s
	s.Width = width
	s.Height = height
	s.World = world
	s.Components.Init(s.World)

	s.Renderables = renderablesSubscription
}

func (s *Scene) InitializeLua() {
	if s.LuaInitialized() {
		return
	}

	s.Lua = lua.NewState()

	for _, luaTypeExporter := range luaTypeExporters {
		luaTypeExport := luaTypeExporter(s)
		common.RegisterLuaType(s.Lua, luaTypeExport)
	}
}

func (s *Scene) UninitializeLua() {
	s.Lua = nil
}

func (s *Scene) LuaInitialized() bool {
	return s.Lua != nil
}

func (s *Scene) updateSceneGraph() {
	for _, eid := range s.Renderables.GetEntities() {
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

func (s *Scene) GenericUpdate(dt time.Duration) {
	s.updateSceneGraph()
	s.updateSceneObjects(dt)
}

func (s *Scene) Render() {
	if len(s.Cameras) < 1 {
		s.initCamera()
	}

	for _, e := range s.Cameras {
		c, found := s.Components.Camera2D.Get(e)
		if !found {
			continue
		}

		rl.BeginMode2D(c.Camera2D)
		s.renderToCamera(e)
		rl.EndMode2D()
	}
}

func (s *Scene) initCamera() {
	s.Cameras = make([]akara.EID, 0)
	s.Cameras = append(s.Cameras, s.Add.Camera(0, 0, s.Width, s.Height))
}

func (s *Scene) renderToCamera(cameraID akara.EID) {
	rt, found := s.Components.RenderTexture2D.Get(cameraID)
	if !found {
		return
	}

	rl.BeginTextureMode(*rt.RenderTexture2D)
	rl.ClearBackground(rl.Black)

	for _, entity := range s.Renderables.GetEntities() {
		if entity == cameraID {
			continue
		}

		s.renderEntity(entity)
	}

	rl.EndTextureMode()
}

func (s *Scene) Key() string {
	return s.key
}
