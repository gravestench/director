package file_loader

import (
	"fmt"
	"io"
	"net/http"
	"os"

	fileLoadRequest "github.com/gravestench/director/pkg/components/file_load_request"
	fileLoadResponse "github.com/gravestench/director/pkg/components/file_load_response"
	fileType "github.com/gravestench/director/pkg/components/file_type"

	"github.com/gravestench/akara"
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
		request  fileLoadRequest.ComponentFactory
		response fileLoadResponse.ComponentFactory
		fileType fileType.ComponentFactory
	}
	subscribed struct {
		awaitingResponse *akara.Subscription
		needsFileType    *akara.Subscription
	}
}

func (sys *System) Name() string {
	return "FileLoader"
}

func (sys *System) Init(world *akara.World) {
	sys.World = world

	sys.initComponents()
	sys.initSubscriptions()
	sys.initProviders()
}

func (sys *System) initComponents() {
	sys.InjectComponent(&fileLoadRequest.Component{}, &sys.components.request.ComponentFactory)
	sys.InjectComponent(&fileLoadResponse.Component{}, &sys.components.response.ComponentFactory)
	sys.InjectComponent(&fileType.Component{}, &sys.components.fileType.ComponentFactory)
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
		Require(&fileLoadRequest.Component{}).
		Forbid(&fileLoadResponse.Component{})

	sys.subscribed.awaitingResponse = sys.AddSubscription(filter.Build())
}

func (sys *System) initTypeResolutionSubscription() {
	filter := sys.World.NewComponentFilter()

	filter.
		Require(&fileLoadResponse.Component{}).
		Forbid(&fileType.Component{})

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
