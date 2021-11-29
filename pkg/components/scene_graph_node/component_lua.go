package scene_graph_node

import lua "github.com/yuin/gopher-lua"

// ExportToLua exports the component into the given table, using the given lua state machine.
func (component *SceneGraphNode) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	// not implemented, but dont want to panic here
	return table
}
