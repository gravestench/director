package origin

import (
	components2 "github.com/gravestench/director/pkg/common/components"
	lua "github.com/yuin/gopher-lua"
)

// ExportToLua exports the component into the given table, using the given lua state machine.
func (component *Origin) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	components2.LuaExportVector3(component.Vector3, state, table)

	return table
}
