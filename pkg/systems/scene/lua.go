package scene

import "github.com/gravestench/director/pkg/common"

var luaTypeExporters = []func(*Scene) common.LuaTypeExport{
	luaRectangleTypeExporter,
	luaCircleTypeExporter,
}
