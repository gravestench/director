package transform

import (
	components2 "github.com/gravestench/director/pkg/common/components"
	lua "github.com/yuin/gopher-lua"
)

const (
	fieldTranslation = "translation"
	fieldRotation    = "rotation"
	fieldScale       = "scale"
)

// ExportToLua exports the component into the given table, using the given lua state machine.
func (component *Transform) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	t := components2.LuaExportVector3(component.Translation, state, state.NewTable())
	r := components2.LuaExportVector3(component.Rotation, state, state.NewTable())
	s := components2.LuaExportVector3(component.Scale, state, state.NewTable())

	state.SetField(table, fieldTranslation, t)
	state.SetField(table, fieldRotation, r)
	state.SetField(table, fieldScale, s)

	return table
}
