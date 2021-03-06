package pkg

import (
	"time"

	"github.com/gravestench/akara"
	"github.com/gravestench/scenegraph"
)

const (
	DefaultSceneTickRate = 60
)

// Scene is the basic scene struct, intended to be embedded in an end-user's scenes.
// This scene "class" has most of the generic scene lifecycle stuff required, so
// it is recommended that it's always used to create new scenes (for now).
type Scene struct {
	SceneSystem
	Graph      scenegraph.Node
	renderList []akara.EID
	Viewports  []akara.EID
	key        string
}

func (s *Scene) Name() string {
	return "BaseScene"
}

func (s *Scene) GenericSceneInit(d *Director) {
	s.Add.scene = s
	s.Director = d
	s.Components.Init(s.Director.World)
	s.SetTickFrequency(DefaultSceneTickRate)
	s.BaseSystem.SetPreTickCallback(func() {
		s.GenericUpdate()
	})
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
	s.updateSceneObjects(s.TimeDelta)

	// this renders the scene objects to the scene's render texture
	// however, this will not actually display anything, that is done by the render system
	s.Render()
}

func (s *Scene) Render() {
	if len(s.Viewports) < 1 {
		s.initViewport()
	}

	for _, e := range s.Viewports {
		s.drawEntitiesAndRender(e)
	}
}

func (s *Scene) RemoveEntity(e akara.EID) {
	for idx := range s.renderList {
		if s.renderList[idx] == e {
			s.renderList = append(s.renderList[:idx], s.renderList[idx+1:]...)
			break
		}
	}

	s.Director.RemoveEntity(e)
}
