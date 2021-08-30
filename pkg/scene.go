package pkg

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/components"
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
	basicComponents
	Graph       scenegraph.Node
	key         string
	Add         SceneObjectFactory
	renderables *akara.Subscription
	Cameras     []akara.EID
}

func (s *Scene) renderEntity(e akara.EID) {
	texture, textureFound := s.Texture2D.Get(e)
	rt, rtFound := s.RenderTexture2D.Get(e)
	if !textureFound && !rtFound {
		return
	}

	var t *rl.Texture2D

	if !rtFound {
		t = &texture.Texture2D
	} else {
		t = &rt.Texture
	}

	trs, found := s.Transform.Get(e)
	if !found {
		return
	}

	x, y := trs.Translation.XY()
	position := rl.Vector2{
		X: float32(x),
		Y: float32(y),
	}

	rotation := float32(trs.Rotation.Y)

	scale := float32(trs.Scale.X)

	rl.DrawTextureEx(*t, position, rotation, scale, rl.White)
}

func (s *Scene) bind(d *Director) {
	s.Director = d
	s.init(d.World)
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
		node, found := s.SceneGraphNode.Get(eid)
		if !found {
			continue
		}

		trs, found := s.Transform.Get(eid)
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
		c, found := s.Camera2D.Get(e)
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
	rt, found := s.RenderTexture2D.Get(cameraID)
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
