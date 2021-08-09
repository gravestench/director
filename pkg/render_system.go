package pkg

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/components"
)

type renderSystem struct {
	akara.BaseSubscriberSystem
	components struct {
		basicComponents
	}
	text *akara.Subscription
}

func (r *renderSystem) Init(world *akara.World) {
	r.World = world

	r.components.init(world)
	r.initTextSubscription()
}

func (r *renderSystem) IsInitialized() bool {
	if r.World == nil {
		return false
	}

	if !r.components.isInit() {
		return false
	}

	return true
}

func (r *renderSystem) initTextSubscription() {
	filter := r.World.NewComponentFilter()

	filter.Require(
		&components.Text{},
		&components.Vector2{},
		&components.Color{},
		)

	r.text = r.World.AddSubscription(filter.Build())
}

func (r *renderSystem) Update() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.Black)

	r.drawText()
}

func (r *renderSystem) drawText() {
	for _, e := range r.text.GetEntities() {
		text, _ := r.components.Text.Get(e)
		vec2, _ := r.components.Vector2.Get(e)
		c, _ := r.components.Color.Get(e)

		rlc := rl.Color{
			R: c.R,
			G: c.G,
			B: c.B,
			A: c.A,
		}

		str := text.String
		x, y := int32(vec2.X), int32(vec2.Y)
		rl.DrawText(str, x, y, 60, rlc)
	}
}

