package pkg

import (
	"github.com/gravestench/director/pkg/common/components"
	lua "github.com/yuin/gopher-lua"
)

const (
	luaSceneTable              = "scene"
	luaSceneSystemsTable       = "sys"        // scene.sys
	luaSceneComponentsTable    = "components" // scene.components
	luaSceneObjectFactoryTable = "add"        // scene.add

	luaConstantsTable = "constants"
)

func (s *Scene) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	panic("implement me")
}

func (s *SceneSystem) initLuaSceneTable() {
	table := s.Lua.NewTable()
	s.Lua.SetGlobal(luaSceneTable, table)

	s.initLuaSceneObjectFactories(table)
	s.initLuaComponentsTable(table)
	s.initLuaSystemsTable(table)
}

func (s *SceneSystem) initLuaConstantsTable() {
	componentsTable := s.Lua.NewTable()
	s.Lua.SetGlobal(luaConstantsTable, componentsTable)

	s.luaExportConstantsInput(componentsTable)
	s.luaExportConstantsLogging(componentsTable)
}

func (s *SceneSystem) initLuaComponentsTable(sceneTable *lua.LTable) {
	componentsTable := s.Lua.NewTable()
	s.Lua.SetField(sceneTable, luaSceneComponentsTable, componentsTable)

	factories := []components.LuaExport{
		&s.Components.Animation,
		&s.Components.Camera,
		&s.Components.Viewport,
		&s.Components.Color,
		&s.Components.Debug,
		&s.Components.FileLoadRequest,
		&s.Components.FileLoadResponse,
		&s.Components.FileType,
		&s.Components.Fill,
		&s.Components.HasChildren,
		&s.Components.Stroke,
		&s.Components.Font,
		&s.Components.Interactive,
		&s.Components.Opacity,
		&s.Components.Origin,
		&s.Components.RenderTexture2D,
		&s.Components.RenderOrder,
		&s.Components.Size,
		&s.Components.SceneGraphNode,
		&s.Components.Text,
		&s.Components.Texture2D,
		&s.Components.Transform,
		&s.Components.UUID,
		&s.Components.Audible,
	}

	for _, factory := range factories {
		factory.ExportToLua(s.Lua, componentsTable)
	}
}

func (s *SceneSystem) initLuaSystemsTable(sceneTable *lua.LTable) {
	sysTable := s.Lua.NewTable()
	s.Lua.SetField(sceneTable, luaSceneSystemsTable, sysTable)

	s.luaExportSystemRenderer(sysTable)
}

func (s *SceneSystem) initLuaSceneObjectFactories(sceneTable *lua.LTable) {
	objFactoryTable := s.Lua.NewTable()

	s.luaBindSceneObjectFactoryCircle(objFactoryTable)
	s.luaBindSceneObjectFactoryImage(objFactoryTable)
	s.luaBindSceneObjectFactoryLabel(objFactoryTable)
	s.luaBindSceneObjectFactoryRectangle(objFactoryTable)

	s.Lua.SetField(sceneTable, luaSceneObjectFactoryTable, objFactoryTable)
}
