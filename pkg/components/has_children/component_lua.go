package has_children

import (
	"strconv"

	"github.com/gravestench/akara"

	lua "github.com/yuin/gopher-lua"
)

const (
	fieldChildren = "children"
)

// ExportToLua exports the component into the given table, using the given lua state machine.
func (component *HasChildren) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	state.SetField(table, fieldChildren, component.luaExportFieldChildren(state, table))

	return table
}

func (component *HasChildren) luaExportFieldChildren(state *lua.LState, table *lua.LTable) *lua.LFunction {
	fnGetAppend := func(L *lua.LState) int {
		numArgs := L.GetTop()

		// yield children in a table
		if numArgs < 1 {
			dataTable := L.NewTable()

			if component.Children == nil {
				return 0
			}

			for idx, child := range component.Children {
				L.SetField(dataTable, strconv.Itoa(idx), lua.LNumber(child))
			}

			L.Push(dataTable)

			return 1
		}

		// else, append children
		list := L.CheckTable(1)
		list.ForEach(func(idx, val lua.LValue) {
			eid, err := strconv.Atoi(val.String())
			if err != nil {
				return
			}

			component.Children = append(component.Children, akara.EID(eid))
		})

		return 0
	}

	return state.NewFunction(fnGetAppend)
}
