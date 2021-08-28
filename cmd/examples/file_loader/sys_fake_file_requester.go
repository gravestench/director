package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gravestench/akara"

	"github.com/gravestench/director/cmd/examples/file_loader/loader"
)

// this would be another system making requests for files

type FakeFileRequester struct {
	*loader.Loader
	akara.BaseSubscriberSystem
	components struct {
		request  LoadRequestFactory
		response FileStreamFactory
	}
	subscribed struct {
		responses *akara.Subscription
	}
	timeUntilNextRequest time.Duration
}

func (sys *FakeFileRequester) Init(world *akara.World) {
	sys.World = world

	sys.initComponents()
	sys.initSubscriptions()
}

func (sys *FakeFileRequester) Update(delta time.Duration) {
	sys.occasionallyMakeRequest(delta)
	sys.checkRequestedFiles()
}

func (sys *FakeFileRequester) occasionallyMakeRequest(delta time.Duration) {
	const (
		loadInterval = time.Millisecond * 3
	)

	sys.timeUntilNextRequest -= delta

	if sys.timeUntilNextRequest.Milliseconds() > 0 {
		return
	}

	sys.timeUntilNextRequest = loadInterval

	if rand.Intn(2) > 0 {
		sys.makeValidRequest()
	} else {
		sys.makeInvalidRequest()
	}
}

func (sys FakeFileRequester) makeValidRequest() {
	requestPath := "testfiles/a.txt"

	fmt.Printf("creating request for file %s\n", requestPath)

	e := sys.World.NewEntity()

	request := sys.components.request.Add(e)

	request.Path = requestPath
}

func (sys FakeFileRequester) makeInvalidRequest() {
	requestPath := "testfiles/nonexistant.txt"

	fmt.Printf("creating request for file %s\n", requestPath)

	e := sys.World.NewEntity()

	request := sys.components.request.Add(e)

	request.Path = requestPath
}

func (sys *FakeFileRequester) checkRequestedFiles() {
	for _, eid := range sys.subscribed.responses.GetEntities() {
		sys.handleResponse(eid)
	}
}

func (sys *FakeFileRequester) handleResponse(e akara.EID) {
	_, found := sys.components.response.Get(e)
	if !found {
		return
	}

	req, found := sys.components.request.Get(e)
	if !found {
		return
	}

	fmt.Printf("Loaded file %s\n", req.Path)

	sys.components.request.Remove(e)
	sys.components.response.Remove(e)
}

func (sys *FakeFileRequester) initComponents() {
	sys.InjectComponent(&FileLoadRequest{}, &sys.components.request.ComponentFactory)
	sys.InjectComponent(&FileLoadResponse{}, &sys.components.response.ComponentFactory)
}

func (sys *FakeFileRequester) initSubscriptions() {
	filter := sys.World.NewComponentFilter().
		Require(&FileLoadResponse{})

	sys.subscribed.responses = sys.AddSubscription(filter.Build())
}

func (sys *FakeFileRequester) IsInitialized() bool {
	ready := sys.components.request.ComponentFactory != nil
	ready = ready && sys.subscribed.responses != nil

	return ready
}

