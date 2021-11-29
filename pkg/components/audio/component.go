package audio

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/audio_source"
	"github.com/gravestench/director/pkg/common/components"
)

// Audio is used to define the state of audio playback for a single source of audio data.
type Audio struct {
	*audio_source.AudioSource
}

// New returns a Audio component.
func (*Audio) New() akara.Component {
	audio := &Audio{}
	audio.AudioSource = audio_source.New()

	return audio
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = Audio // Component is an alias to Audio
