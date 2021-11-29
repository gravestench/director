package scene_graph_node

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common/components"
	"github.com/gravestench/scenegraph"
)

var _ akara.Component = &SceneGraphNode{}

// SceneGraphNode is a component that contains normalized alpha transparency (0.0 ... 1.0)
type SceneGraphNode struct {
	*scenegraph.Node
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*SceneGraphNode) New() akara.Component {
	return &SceneGraphNode{
		Node: scenegraph.NewNode(),
	}
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = SceneGraphNode // Component is an alias to SceneGraphNode
