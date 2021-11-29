package file_load_response

import (
	"io"
	"strconv"

	lua "github.com/yuin/gopher-lua"
)

// ExportToLua exports the component into the given table, using the given lua state machine.
func (res *FileLoadResponse) ExportToLua(state *lua.LState, table *lua.LTable) *lua.LTable {
	fnGetDataTable := func(L *lua.LState) int {
		if L.GetTop() != 0 {
			return 0
		}

		dataTable := L.NewTable()

		if res.Stream == nil {
			return 0
		}

		_, err := res.Stream.Seek(0, io.SeekStart)
		if err != nil {
			return 0
		}

		bytes, err := io.ReadAll(res.Stream)
		if err != nil {
			return 0
		}

		for idx := range bytes {
			L.SetField(dataTable, strconv.Itoa(idx), lua.LNumber(bytes[idx]))
		}

		L.Push(dataTable)

		return 1
	}

	state.SetField(table, "data", state.NewFunction(fnGetDataTable))

	return table
}
