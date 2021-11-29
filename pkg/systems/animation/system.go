package animation

import (
	"time"

	texture2D "github.com/gravestench/director/pkg/components/texture"

	"github.com/gravestench/director/pkg/components/animation"

	"github.com/gravestench/akara"
)

type System struct {
	akara.BaseSystem
	components struct {
		animations animation.ComponentFactory
		textures   texture2D.ComponentFactory
	}
	animations *akara.Subscription
}

func (sys *System) Init(_ *akara.World) {
	sys.initComponents()
	sys.initSubscriptions()
}

func (sys *System) Name() string {
	return "Animation"
}

func (sys *System) initComponents() {
	sys.InjectComponent(&texture2D.Component{}, &sys.components.textures.ComponentFactory)
	sys.InjectComponent(&animation.Component{}, &sys.components.animations.ComponentFactory)
}

func (sys *System) initSubscriptions() {
	filter := sys.NewComponentFilter()

	filter.Require(&texture2D.Component{})
	filter.Require(&animation.Animation{})

	sys.animations = sys.AddSubscription(filter.Build())
}

func (sys *System) IsInitialized() bool {
	return sys.animations != nil
}

func (sys *System) Update() {
	for _, e := range sys.animations.GetEntities() {
		sys.updateAnimation(e, sys.TimeDelta)
	}
}

func (sys *System) updateAnimation(e akara.EID, dt time.Duration) {
	anim, found := sys.components.animations.Get(e)
	if !found {
		return
	}

	tex, found := sys.components.textures.Get(e)
	if !found {
		return
	}

	anim.UntilNextFrame -= dt

	for anim.UntilNextFrame <= 0 {
		anim.CurrentFrame += 1
		anim.CurrentFrame %= len(anim.FrameImages)
		anim.UntilNextFrame += anim.FrameDurations[anim.CurrentFrame]
	}

	tex.Texture2D = anim.FrameTextures[anim.CurrentFrame]
}
