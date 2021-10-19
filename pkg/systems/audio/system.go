package audio

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/components"
	"io"
)

const (
	SampleRate        = 48000
	BufferSizeSamples = 2400
)

// static check that System implements the System interface
var _ akara.System = &System{}

// System is responsible for handling Audible entities
type System struct {
	akara.BaseSystem
	subscriptions struct {
		needsFile *akara.Subscription
		sounds    *akara.Subscription
	}
	initialized bool
	Components  struct {
		Audible          AudibleFactory
		fileType         components.FileTypeFactory
		fileLoadResponse components.FileLoadResponseFactory
	}
}

func (m *System) Name() string {
	return "Audio"
}

func (m *System) IsInitialized() bool {
	return m.initialized
}

// Init initializes the system with the given world, injecting the necessary components
func (m *System) Init(_ *akara.World) {
	m.setupFactories()
	m.setupSubscriptions()

	err := speaker.Init(SampleRate, BufferSizeSamples)
	if err != nil {
		panic("Failed to initialize audio system: " + err.Error())
	}

	m.initialized = true
}

func (m *System) setupFactories() {
	m.InjectComponent(&Audible{}, &m.Components.Audible.ComponentFactory)
	m.InjectComponent(&components.FileLoadResponse{}, &m.Components.fileLoadResponse.ComponentFactory)
	m.InjectComponent(&components.FileType{}, &m.Components.fileType.ComponentFactory)
}

func (m *System) setupSubscriptions() {
	sounds := m.NewComponentFilter().
		Require(&Audible{}).
		Build()

	m.subscriptions.sounds = m.AddSubscription(sounds)

	needsFile := m.NewComponentFilter().
		Require(&Audible{}).
		Require(&components.FileType{}).
		Require(&components.FileLoadResponse{}).
		Build()

	m.subscriptions.needsFile = m.AddSubscription(needsFile)
}

// Update will iterate over interactive entities
func (m *System) Update() {
	for _, id := range m.subscriptions.needsFile.GetEntities() {
		m.addStreamToAudible(id)
	}
}

func (m *System) addStreamToAudible(id akara.EID) {
	audible, _ := m.Components.Audible.Get(id)
	fileStream, _ := m.Components.fileLoadResponse.Get(id)
	fileType, _ := m.Components.fileType.Get(id)

	var stream beep.StreamSeekCloser
	var format beep.Format
	var err error
	switch fileType.Type {
	case "audio/wave":
		_, _ = fileStream.Stream.Seek(0, io.SeekStart)
		stream, format, err = wav.Decode(fileStream.Stream)
	default:
		fmt.Printf("Unsupported audio file format: %s\n", fileType.Type)
		return
	}

	if err != nil {
		// TODO: logging system? something nicer than just printing?
		fmt.Printf("Failed to decode file: %s\n", err.Error())
		return
	}

	audible.setStream(stream, format)

	// we've handled this fileLoadResponse, so we don't need to keep seeing it in the subscription anymore
	m.subscriptions.needsFile.IgnoreEntity(id)
}
