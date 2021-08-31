package pkg

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/components"
	"github.com/gravestench/mathlib"
	"github.com/gravestench/scenegraph"
	"time"
)

type SceneFace interface {
	akara.System
	bindsToDirector
	hasKey
	update(duration time.Duration)
	render()
}

type bindsToDirector interface {
	bind(*Director)
	unbind()
}

type hasKey interface {
	Key() string
}

type Scene struct {
	*Director
	akara.BaseSystem
	Components basicComponents
	Graph       scenegraph.Node
	key         string
	Add         SceneObjectFactory
	renderables *akara.Subscription
	Cameras     []akara.EID
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

func (s *Scene) bind(d *Director) {
	s.Director = d
	s.Components.init(d.World)
	s.Add.scene = s

	s.initRenderablesSubscription()
}

func (s *Scene) initRenderablesSubscription() {
	f := s.Director.NewComponentFilter()

	f.Require(&components.SceneGraphNode{})
	f.Require(&components.Transform{})
	f.RequireOne(&components.RenderTexture2D{}, &components.Texture2D{})

	s.renderables = s.Director.AddSubscription(f.Build())
}

func (s *Scene) unbind() {
	s.Director = nil
}

func (s *Scene) updateSceneGraph() {
	for _, eid := range s.renderables.GetEntities() {
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

func (s *Scene) update(dt time.Duration) {
	s.updateSceneGraph()
	s.updateSceneObjects(dt)
}

func (s *Scene) render() {
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
	w, h := s.Director.Window.Width, s.Director.Window.Height
	s.Cameras = make([]akara.EID, 0)
	s.Cameras = append(s.Cameras, s.Add.Camera(0, 0, w, h))
}

func (s *Scene) renderToCamera(cameraID akara.EID) {
	rt, found := s.Components.RenderTexture2D.Get(cameraID)
	if !found {
		return
	}

	rl.BeginTextureMode(*rt.RenderTexture2D)
	rl.ClearBackground(rl.Black)

	for _, entity := range s.renderables.GetEntities() {
		if entity == cameraID {
			continue
		}

		s.renderEntity(entity)
	}

	rl.EndTextureMode()
}
