package audio

import (
	"fmt"
	"io"

	"github.com/gravestench/director/pkg/audio_source"

	fileLoadResponse "github.com/gravestench/director/pkg/components/file_load_response"

	fileType "github.com/gravestench/director/pkg/components/file_type"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/components/audio"
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
		Audible          audio.ComponentFactory
		fileType         fileType.ComponentFactory
		fileLoadResponse fileLoadResponse.ComponentFactory
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

	err := speaker.Init(audio_source.SampleRate, audio_source.BufferSizeSamples)
	if err != nil {
		panic("Failed to initialize audio system: " + err.Error())
	}

	m.initialized = true
}

func (m *System) setupFactories() {
	m.InjectComponent(&audio.Audio{}, &m.Components.Audible.ComponentFactory)
	m.InjectComponent(&fileLoadResponse.Component{}, &m.Components.fileLoadResponse.ComponentFactory)
	m.InjectComponent(&fileType.Component{}, &m.Components.fileType.ComponentFactory)
}

func (m *System) setupSubscriptions() {
	sounds := m.NewComponentFilter().
		Require(&audio.Audio{}).
		Build()

	m.subscriptions.sounds = m.AddSubscription(sounds)

	needsFile := m.NewComponentFilter().
		Require(&audio.Audio{}).
		Require(&fileType.Component{}).
		Require(&fileLoadResponse.Component{}).
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

	audible.AudioSource.SetStream(stream, format)

	// we've handled this fileLoadResponse, so we don't need to keep seeing it in the subscription anymore
	m.subscriptions.needsFile.IgnoreEntity(id)
}
