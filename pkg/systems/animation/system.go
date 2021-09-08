package animation

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/components"
	"time"
)

type System struct {
	akara.BaseSystem
	components struct {
		animations components.AnimationFactory
		textures   components.Texture2DFactory
	}
	animations *akara.Subscription
}

func (sys *System) Init(_ *akara.World) {
	sys.initComponents()
	sys.initSubscriptions()
}

func (sys *System) initComponents() {
	sys.InjectComponent(&components.Texture2D{}, &sys.components.textures.ComponentFactory)
	sys.InjectComponent(&components.Animation{}, &sys.components.animations.ComponentFactory)
}

func (sys *System) initSubscriptions() {
	filter := sys.NewComponentFilter()

	filter.Require(&components.Texture2D{})
	filter.Require(&components.Animation{})

	sys.animations = sys.AddSubscription(filter.Build())
}

func (sys *System) IsInitialized() bool {
	return sys.animations != nil
}

func (sys *System) Update(dt time.Duration) {
	for _, e := range sys.animations.GetEntities() {
		sys.updateAnimation(e, dt)
	}
}

func (sys *System) updateAnimation(e common.Entity, dt time.Duration) {
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

