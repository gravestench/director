package scene

//const (
//	luaComponentsName = "components"
//)
//
//// Registers my transform type to given L.
//func registerTransformType(L *lua.LState) {
//	mt := L.NewTypeMetatable(luaComponentsName)
//	tblComponents := L.NewTable()
//
//	L.SetGlobal(luaComponentsName, mt)
//	// static attributes
//	L.SetField(tblComponents, "Transform", L.NewFunction(newTransform))
//	L.SetField(mt, "new", L.NewFunction(newTransform))
//	// methods
//	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), transformMethods))
//}
//
// Constructor
//func newTransform(L *lua.LState) int {
//	transform := &components.Transform{
//		Translation: &mathlib.Vector3{
//			X: float64(L.CheckNumber(1)),
//			Y: float64(L.CheckNumber(2)),
//			Z: float64(L.CheckNumber(3)),
//		},
//	}
//	ud := L.NewUserData()
//	ud.Value = transform
//	L.SetMetatable(ud, L.GetTypeMetatable(luaComponentsName))
//	L.Push(ud)
//	return 1
//}
//
//// Checks whether the first lua argument is a *LUserData with *Transform and returns this *Transform.
//func checkTransform(L *lua.LState) *components.Transform {
//	ud := L.CheckUserData(1)
//	if v, ok := ud.Value.(*components.Transform); ok {
//		return v
//	}
//	L.ArgError(1, "transform expected")
//	return nil
//}
//
//var transformMethods = map[string]lua.LGFunction{
//	"translation": transformGetSetTranslation,
//}
//
//// Getter and setter for the Transform#Name
//func transformGetSetTranslation(L *lua.LState) int {
//	p := checkTransform(L)
//	if L.GetTop() == 3 {
//		p.Name = L.CheckString(2)
//		return 0
//	}
//
//	L.Push(lua.LString(p.Name))
//	return 1
//}
