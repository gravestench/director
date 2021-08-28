package main

import (
	"fmt"
	"github.com/gravestench/akara"
	"github.com/gravestench/director/cmd/examples/file_loader/loader"
)

const allowedFailedLoadAttempts = 10

type FileLoader struct {
	*loader.Loader
	akara.BaseSubscriberSystem
	components struct {
		request  LoadRequestFactory
		response FileStreamFactory
	}
	subscribed struct {
		requests *akara.Subscription
	}
}

func (sys *FileLoader) Update() {
	for _, eid := range sys.subscribed.requests.GetEntities() {
		sys.processLoadRequest(eid)
	}
}

func (sys *FileLoader) processLoadRequest(e akara.EID) {
	requested, found := sys.components.request.Get(e)
	if !found {
		return
	}

	requested.Attempts++
	if requested.Attempts >= allowedFailedLoadAttempts {
		fmt.Printf("Could not find %s after %v attempts\n", requested.Path, requested.Attempts)
		sys.components.request.Remove(e)

		return
	}

	stream, err := sys.Load(requested.Path)
	if err != nil {
		return
	}

	fs := sys.components.response.Add(e)
	fs.ReadSeeker = stream
}

func (sys *FileLoader) Init(world *akara.World) {
	sys.World = world

	sys.initLoader()
	sys.initComponents()
	sys.initSubscriptions()
}

func (sys *FileLoader) initLoader() {
	sys.Loader = loader.New()
}

func (sys *FileLoader) initComponents() {
	sys.InjectComponent(&FileLoadRequest{}, &sys.components.request.ComponentFactory)
	sys.InjectComponent(&FileLoadResponse{}, &sys.components.response.ComponentFactory)
}

func (sys *FileLoader) initSubscriptions() {
	sys.initLoadRequestSubscription()
}

func (sys *FileLoader) initLoadRequestSubscription() {
	filter := sys.World.NewComponentFilter()

	filter.Require(&FileLoadRequest{})
	filter.Forbid(&FileLoadResponse{})

	sys.subscribed.requests = sys.AddSubscription(filter.Build())
}

func (sys *FileLoader) IsInitialized() bool {
	ready := sys.Loader != nil
	ready = ready && sys.components.request.ComponentFactory != nil
	ready = ready && sys.subscribed.requests != nil

	return ready
}

