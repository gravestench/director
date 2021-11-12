package pkg

import (
	lua "github.com/yuin/gopher-lua"
)

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

	s.luaExportComponentAnimation(componentsTable)
	s.luaExportComponentCamera(componentsTable)
	s.luaExportComponentColor(componentsTable)
	s.luaExportComponentDebug(componentsTable)
	s.luaExportComponentFileLoadRequest(componentsTable)
	s.luaExportComponentFileLoadResponse(componentsTable)
	// s.luaExportComponentFileType(componentsTable)
	s.luaExportComponentFill(componentsTable)
	s.luaExportComponentFont(componentsTable)
	s.luaExportComponentHasChildren(componentsTable)
	s.luaExportComponentInteractive(componentsTable)
	s.luaExportComponentOpacity(componentsTable)
	s.luaExportComponentOrigin(componentsTable)
	s.luaExportComponentRenderOrder(componentsTable)
	// s.luaExportComponentRender_texture(componentsTable)
	// s.luaExportComponentScene_graph_node(componentsTable)
	s.luaExportComponentSize(componentsTable)
	s.luaExportComponentStroke(componentsTable)
	s.luaExportComponentText(componentsTable)
	// s.luaExportComponentTexture(componentsTable)
	s.luaExportComponentTransform(componentsTable)
	s.luaExportComponentUUID(componentsTable)
	// s.luaExportComponentViewport(componentsTable)
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
