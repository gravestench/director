package scene

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	director "github.com/gravestench/director/pkg"
	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/components"
	"github.com/gravestench/mathlib"
	"github.com/gravestench/scenegraph"
	lua "github.com/yuin/gopher-lua"
	"reflect"
	"strings"
	"time"
)

type Scene struct {
	akara.BaseSystem
	*director.Director
	Lua         *lua.LState
	Components  common.BasicComponents
	Graph       scenegraph.Node
	key         string
	Add         ObjectFactory
	Renderables *akara.Subscription
	Cameras     []akara.EID
	Width       int
	Height      int
}

var tmpVect mathlib.Vector3

func (s *Scene) renderEntity(e akara.EID) {
	texture, textureFound := s.Components.Texture2D.Get(e)
	rt, rtFound := s.Components.RenderTexture2D.Get(e)
	if !textureFound && !rtFound {
		return
	}

	var t *rl.Texture2D

	if !rtFound {
		t = &texture.Texture2D
	} else {
		t = &rt.Texture
	}

	trs, found := s.Components.Transform.Get(e)
	if !found {
		return
	}

	origin, found := s.Components.Origin.Get(e)
	if !found {
		return
	}

	tmpVect.Set(float64(rt.Texture.Width), float64(rt.Texture.Height), 1)

	yRad := trs.Rotation.Y * mathlib.DegreesToRadians
	ov2 := mathlib.NewVector2(origin.Clone().Multiply(&tmpVect).XY()).Rotate(yRad).Negate()
	ov3 := mathlib.NewVector3(ov2.X, ov2.Y, 0)

	x, y := trs.Translation.Clone().Add(ov3).XY()
	v2 := mathlib.NewVector2(x, y)

	position := rl.Vector2{
		X: float32(v2.X),
		Y: float32(v2.Y),
	}

	rotation := float32(trs.Rotation.Y)

	scale := float32(trs.Scale.X)

	rl.DrawTextureEx(*t, position, rotation, scale, rl.White)
}

func (s *Scene) Initialize(d *director.Director, width, height int) {
	s.Add.scene = s
	s.Width = width
	s.Height = height
	s.Director = d
	s.Components.Init(s.Director.World)

	filter := s.Director.World.NewComponentFilter()
	filter.Require(&components.Transform{})
	filter.RequireOne(&components.RenderTexture2D{}, &components.Texture2D{})

	s.Renderables = s.Director.AddSubscription(filter.Build())
}

func (s *Scene) InitializeLua() {
	if s.LuaInitialized() {
		return
	}

	s.Lua = lua.NewState()

	for _, luaTypeExporter := range luaTypeExporters {
		luaTypeExport := luaTypeExporter(s)
		common.RegisterLuaType(s.Lua, luaTypeExport)
	}

	s.initComponentsTable()
}

func (s *Scene) initComponentsTable() {
	componentsTable := s.Lua.NewTable()
	s.Lua.SetGlobal("components", componentsTable)

	s.addTransformComponent(componentsTable)
}

func (s *Scene) addTransformComponent(mt *lua.LTable) {
	name := strings.ToLower(reflect.TypeOf(&components.Transform{}).Elem().Name())
	trsTable := s.Lua.NewTypeMetatable(name)

	checkEid := func() *akara.EID {
		ud := s.Lua.CheckUserData(1)
		if v, ok := ud.Value.(*akara.EID); ok {
			return v
		}
		s.Lua.ArgError(1, "EID expected")
		return nil
	}

	makeVec3Table := func(vec3 *mathlib.Vector3) *lua.LFunction {
		setGetXYZ := func(L *lua.LState) int {
			if L.GetTop() == 3 {
				x := L.CheckNumber(1)
				y := L.CheckNumber(2)
				z := L.CheckNumber(3)

				vec3.Set(float64(x), float64(y), float64(z))

				return 0
			}

			x, y, z := vec3.XYZ()

			s.Lua.Push(lua.LNumber(x))
			s.Lua.Push(lua.LNumber(y))
			s.Lua.Push(lua.LNumber(z))

			return 3
		}

		return s.Lua.NewFunction(setGetXYZ)
	}

	trsAdd := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		trs := s.Components.Transform.Add(*checkEid())

		table := L.NewTable()
		L.SetField(table, "translation", makeVec3Table(trs.Translation))
		L.SetField(table, "rotation", makeVec3Table(trs.Rotation))
		L.SetField(table, "scale", makeVec3Table(trs.Scale))

		ud := L.NewUserData()
		ud.Value = table
		L.SetMetatable(ud, L.GetTypeMetatable(name))
		L.Push(ud)
		return 1
	}

	trsGet := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		id := L.CheckNumber(1)
		trs, found := s.Components.Transform.Get(akara.EID(id))

		table := L.NewTable()
		L.SetField(table, "translation", makeVec3Table(trs.Translation))
		L.SetField(table, "rotation", makeVec3Table(trs.Rotation))
		L.SetField(table, "scale", makeVec3Table(trs.Scale))

		truthy := lua.LFalse

		if found {
			truthy = lua.LTrue
		}

		L.SetMetatable(table, L.GetTypeMetatable(name))

		L.Push(table)
		L.Push(truthy)

		return 2
	}

	s.Lua.SetField(trsTable, "add", s.Lua.NewFunction(trsAdd))
	s.Lua.SetField(trsTable, "get", s.Lua.NewFunction(trsGet))

	s.Lua.SetField(mt, name, trsTable)
}

func (s *Scene) UninitializeLua() {
	s.Lua = nil
}

func (s *Scene) LuaInitialized() bool {
	return s.Lua != nil
}

func (s *Scene) updateSceneGraph() {
	for _, eid := range s.Renderables.GetEntities() {
		node, found := s.Components.SceneGraphNode.Get(eid)
		if !found {
			continue
		}

		trs, found := s.Components.Transform.Get(eid)
		if !found {
			continue
		}

		node.Local = trs.GetMatrix()
	}

	s.Graph.UpdateWorldMatrix()
}

func (s *Scene) updateSceneObjects(dt time.Duration) {
	s.Add.update(dt)
}

func (s *Scene) GenericUpdate(dt time.Duration) {
	s.updateSceneGraph()
	s.updateSceneObjects(dt)
}

func (s *Scene) Render() {
	if len(s.Cameras) < 1 {
		s.initCamera()
	}

	for _, e := range s.Cameras {
		c, found := s.Components.Camera2D.Get(e)
		if !found {
			continue
		}

		rl.BeginMode2D(c.Camera2D)
		s.renderToCamera(e)
		rl.EndMode2D()
	}
}

func (s *Scene) initCamera() {
	s.Cameras = make([]akara.EID, 0)
	s.Cameras = append(s.Cameras, s.Add.Camera(0, 0, s.Width, s.Height))
}

func (s *Scene) renderToCamera(cameraID akara.EID) {
	rt, found := s.Components.RenderTexture2D.Get(cameraID)
	if !found {
		return
	}

	cam, found := s.Components.Camera2D.Get(cameraID)
	if !found {
		return
	}

	rl.BeginTextureMode(*rt.RenderTexture2D)
	r, g, b, a := cam.Background.RGBA()
	rl.ClearBackground(rl.NewColor(uint8(r), uint8(g), uint8(b), uint8(a)))

	for _, entity := range s.Renderables.GetEntities() {
		if entity == cameraID {
			continue
		}

		s.renderEntity(entity)
	}

	rl.EndTextureMode()
}

func (s *Scene) Key() string {
	return s.key
}
