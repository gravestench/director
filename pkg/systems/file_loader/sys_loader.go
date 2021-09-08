package file_loader

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/components"
	"github.com/gravestench/director/pkg/systems/file_loader/loader"
	"github.com/gravestench/director/pkg/systems/file_loader/loader/file_system"
	"github.com/gravestench/director/pkg/systems/file_loader/loader/web"
)

const allowedFailedLoadAttempts = 10

var _ akara.System = &System{}

type System struct {
	akara.BaseSystem
	loader.Loader
	components struct {
		request  components.FileLoadRequestFactory
		response components.FileLoadResponseFactory
		fileType components.FileTypeFactory
	}
	subscribed struct {
		awaitingResponse *akara.Subscription
		needsFileType    *akara.Subscription
	}
}

func (sys *System) Init(world *akara.World) {
	sys.World = world

	sys.initComponents()
	sys.initSubscriptions()
	sys.initProviders()
}

func (sys *System) initComponents() {
	sys.InjectComponent(&components.FileLoadRequest{}, &sys.components.request.ComponentFactory)
	sys.InjectComponent(&components.FileLoadResponse{}, &sys.components.response.ComponentFactory)
	sys.InjectComponent(&components.FileType{}, &sys.components.fileType.ComponentFactory)
}

func (sys *System) initProviders() {
	cwd, err := os.Getwd()
	if err == nil {
		fsProvider := file_system.New(cwd)
		sys.AddProvider(fsProvider)
	}

	sys.AddProvider(web.New())
}


func (sys *System) initSubscriptions() {
	sys.initLoadRequestSubscription()
	sys.initTypeResolutionSubscription()
}

func (sys *System) initLoadRequestSubscription() {
	filter := sys.World.NewComponentFilter()

	filter.
		Require(&components.FileLoadRequest{}).
		Forbid(&components.FileLoadResponse{})

	sys.subscribed.awaitingResponse = sys.AddSubscription(filter.Build())
}

func (sys *System) initTypeResolutionSubscription() {
	filter := sys.World.NewComponentFilter()

	filter.
		Require(&components.FileLoadResponse{}).
		Forbid(&components.FileType{})

	sys.subscribed.needsFileType = sys.AddSubscription(filter.Build())
}

func (sys *System) IsInitialized() bool {
	if sys.components.request.ComponentFactory == nil {
		return false
	}

	if sys.subscribed.awaitingResponse == nil {
		return false
	}

	return true
}

func (sys *System) Update() {
	for _, eid := range sys.subscribed.awaitingResponse.GetEntities() {
		sys.processLoadRequest(eid)
	}

	for _, eid := range sys.subscribed.needsFileType.GetEntities() {
		sys.determineFileType(eid)
	}
}

func (sys *System) processLoadRequest(e akara.EID) {
	req, found := sys.components.request.Get(e)
	if !found {
		return
	}

	req.Attempts++

	stream, streamErr := sys.Load(req.Path)
	if streamErr != nil {
		if req.Attempts >= allowedFailedLoadAttempts {
			fmt.Printf("Could not find %s after %v attempts\n", req.Path, req.Attempts)
			sys.RemoveEntity(e)
		}

		return
	}

	res := sys.components.response.Add(e)
	res.Stream = stream
}

func (sys *System) determineFileType(e akara.EID) {
	res, found := sys.components.response.Get(e)
	if !found {
		return
	}

	if res.Stream == nil {
		return
	}

	data, err := io.ReadAll(res.Stream)
	if err != nil {
		return
	}

	ft := sys.components.fileType.Add(e)
	ft.Type = http.DetectContentType(data)
}
