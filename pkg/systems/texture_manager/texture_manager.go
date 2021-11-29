package texture_manager

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"time"

	fileLoadRequest "github.com/gravestench/director/pkg/components/file_load_request"

	fileLoadResponse "github.com/gravestench/director/pkg/components/file_load_response"
	fileType "github.com/gravestench/director/pkg/components/file_type"
	texture2D "github.com/gravestench/director/pkg/components/texture"

	"github.com/gravestench/director/pkg/components/animation"

	"github.com/faiface/mainthread"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common/cache"
)

const (
	textureBudget = 10000
)

type System struct {
	akara.BaseSystem
	*cache.Cache
	components struct {
		fileLoadRequest  fileLoadRequest.ComponentFactory
		fileLoadResponse fileLoadResponse.ComponentFactory
		fileType         fileType.ComponentFactory
		texture2d        texture2D.ComponentFactory
		animations       animation.ComponentFactory
	}
	subscriptions struct {
		needsTexture *akara.Subscription
	}
}

func (sys *System) Name() string {
	return "TextureManager"
}

func (sys *System) Init(world *akara.World) {
	sys.World = world

	sys.Cache = cache.New(textureBudget)

	sys.initComponents()
	sys.initSubscriptions()
}

func (sys *System) initComponents() {
	sys.InjectComponent(&fileLoadRequest.Component{}, &sys.components.fileLoadRequest.ComponentFactory)
	sys.InjectComponent(&fileLoadResponse.Component{}, &sys.components.fileLoadResponse.ComponentFactory)
	sys.InjectComponent(&fileType.Component{}, &sys.components.fileType.ComponentFactory)
	sys.InjectComponent(&texture2D.Component{}, &sys.components.texture2d.ComponentFactory)
	sys.InjectComponent(&animation.Animation{}, &sys.components.animations.ComponentFactory)
}

func (sys *System) initSubscriptions() {
	sys.initTextureQueueSubscription()
}

func (sys *System) initTextureQueueSubscription() {
	filter := sys.World.NewComponentFilter()

	filter.
		Require(&fileLoadResponse.Component{}).
		Require(&fileType.Component{}).
		Forbid(&texture2D.Component{})

	sys.subscriptions.needsTexture = sys.AddSubscription(filter.Build())
}

func (sys *System) IsInitialized() bool {
	if sys.components.fileType.ComponentFactory == nil {
		return false
	}

	if sys.subscriptions.needsTexture == nil {
		return false
	}

	return true
}

func (sys *System) Update() {
	for _, e := range sys.subscriptions.needsTexture.GetEntities() {
		sys.createTexture(e)
	}
}

func (sys *System) createTextureFromImage(e akara.EID, img image.Image) {
	req, found := sys.components.fileLoadRequest.Get(e)
	if !found {
		return
	}

	mainthread.Call(func() {
		texture := rl.LoadTextureFromImage(rl.NewImageFromImage(&imageBugHack{img: img}))

		_ = sys.Cache.Insert(req.Path, &texture, 1)

		t := sys.components.texture2d.Add(e)
		t.Texture2D = &texture
	})
}

func (sys *System) NewTextureFromImage(img image.Image) akara.EID {
	e := sys.NewEntity()

	sys.createTextureFromImage(e, img)

	return e
}

func (sys *System) createTexture(e akara.EID) {
	req, found := sys.components.fileLoadRequest.Get(e)
	if !found {
		return
	}

	ft, found := sys.components.fileType.Get(e)
	if !found {
		return
	}

	res, found := sys.components.fileLoadResponse.Get(e)
	if !found {
		return
	}

	if entry, found := sys.Cache.Retrieve(req.Path); found {
		sys.components.texture2d.Add(e).Texture2D = entry.(*rl.Texture2D)
		return
	}

	var img image.Image
	var err error
	switch ft.Type {
	case "image/png":
		_, _ = res.Stream.Seek(0, io.SeekStart)
		img, err = png.Decode(res.Stream)
	case "image/jpg", "image/jpeg":
		_, _ = res.Stream.Seek(0, io.SeekStart)
		img, err = jpeg.Decode(res.Stream)
	case "image/gif":
		_, _ = res.Stream.Seek(0, io.SeekStart)
		var gifImage *gif.GIF

		gifImage, err = gif.DecodeAll(res.Stream)

		if gifImage != nil {
			if len(gifImage.Image) > 1 {
				sys.createGifAnimation(e, gifImage)
				return // we handle cache stuff inside of createGifAnimation
			}

			img = gifImage.Image[0]
		}
	default:
		// we don't handle this file type, ignore this entity
		sys.subscriptions.needsTexture.IgnoreEntity(e)
		return
	}

	if img == nil || err != nil {
		return
	}

	sys.createTextureFromImage(e, img)

	// we've successfully created the texture for this image, so we can ignore this entity ID now
	sys.subscriptions.needsTexture.IgnoreEntity(e)
}

func (sys *System) createGifAnimation(e akara.EID, gifImg *gif.GIF) {
	req, found := sys.components.fileLoadRequest.Get(e)
	if !found {
		return
	}

	anim := sys.components.animations.Add(e)

	mainthread.Call(func() {
		for idx := range gifImg.Image {
			anim.FrameImages = append(anim.FrameImages, gifImg.Image[idx])
			cacheKey := fmt.Sprintf("%s::frame%v", req.Path, idx)

			delay := time.Second / 100 * time.Duration(gifImg.Delay[idx])
			anim.FrameDurations = append(anim.FrameDurations, delay)

			if t, found := sys.Cache.Retrieve(cacheKey); found {
				anim.FrameTextures = append(anim.FrameTextures, t.(*rl.Texture2D))
				continue
			}

			texture := rl.LoadTextureFromImage(rl.NewImageFromImage(&imageBugHack{img: anim.FrameImages[idx]}))
			anim.FrameTextures = append(anim.FrameTextures, &texture)

			_ = sys.Cache.Insert(cacheKey, &texture, 1)
		}

		anim.UntilNextFrame = anim.FrameDurations[0]

		t := sys.components.texture2d.Add(e)
		t.Texture2D = anim.FrameTextures[0]
	})

}
