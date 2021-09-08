package texture_manager

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common/cache"
	"github.com/gravestench/director/pkg/components"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

const (
	textureBudget = 1000
)

type System struct {
	akara.BaseSystem
	*cache.Cache
	components struct {
		fileLoadRequest components.FileLoadRequestFactory
		fileLoadResponse components.FileLoadResponseFactory
		fileType components.FileTypeFactory
		texture2d components.Texture2DFactory
	}
	subscriptions struct {
		needsTexture *akara.Subscription
	}
}

func (sys *System) Init(world *akara.World) {
	sys.World = world

	sys.Cache = cache.New(textureBudget)

	sys.initComponents()
	sys.initSubscriptions()
}

func (sys *System) initComponents() {
	sys.InjectComponent(&components.FileLoadRequest{}, &sys.components.fileLoadRequest.ComponentFactory)
	sys.InjectComponent(&components.FileLoadResponse{}, &sys.components.fileLoadResponse.ComponentFactory)
	sys.InjectComponent(&components.FileType{}, &sys.components.fileType.ComponentFactory)
	sys.InjectComponent(&components.Texture2D{}, &sys.components.texture2d.ComponentFactory)
}

func (sys *System) initSubscriptions() {
	sys.initTextureQueueSubscription()
}

func (sys *System) initTextureQueueSubscription() {
	filter := sys.World.NewComponentFilter()

	filter.
		Require(&components.FileLoadResponse{}).
		Require(&components.FileType{}).
		Forbid(&components.Texture2D{})

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

func (sys *System) createTexture(e akara.EID) {
	var img image.Image

	var err error

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
		sys.components.texture2d.Add(e).Texture2D = entry.(rl.Texture2D)
		return
	}

	_, _ = res.Stream.Seek(0, io.SeekStart)

	switch ft.Type {
	case "image/png":
		img, err = png.Decode(res.Stream)
	case "image/jpg", "image/jpeg":
		img, err = jpeg.Decode(res.Stream)
	default:
		return
	}

	if err != nil {
		return
	}

	t := sys.components.texture2d.Add(e)
	texture := rl.LoadTextureFromImage(rl.NewImageFromImage(&imageBugHack{img: img}))

	_ = sys.Cache.Insert(req.Path, texture, 1)

	t.Texture2D = texture
}

