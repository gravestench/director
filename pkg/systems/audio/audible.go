package audio

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/speaker"
	"github.com/gravestench/akara"
	"time"
)

// static check that Audible implements Component
var _ akara.Component = &Audible{}

// Audible is used to define the state of audio playback for a single source of audio data.
type Audible struct {
	// active denotes whether the speaker holds a reference to this Audible
	active           bool
	looping          bool
	rawDataStream    beep.StreamSeekCloser
	rawDataFormat    beep.Format
	controller       *beep.Ctrl
	volumeController *effects.Volume
	panController    *effects.Pan
	topLevelStream   beep.Streamer

	// FinishedCallback is called when the Audible has finished playing. This is never called if the Audible is looping.
	// Receives one argument, the Audible itself.
	finishedCallback func(*Audible)
}

// New returns a Audible component.
func (*Audible) New() akara.Component {
	audible := &Audible{}
	audible.controller = &beep.Ctrl{Streamer: audible.rawDataStream}
	audible.volumeController = &effects.Volume{
		Streamer: audible.controller,
		Base:     2,
	}
	audible.panController = &effects.Pan{
		Streamer: audible.volumeController,
	}

	audible.topLevelStream = audible.panController

	return audible
}

func (a *Audible) setStream(stream beep.StreamSeekCloser, format beep.Format) {
	a.rawDataStream = stream
	a.rawDataFormat = format

	a.refreshEffectsChain()

	// immediately play the Audible if it was not created in the Paused state
	if !a.Paused() {
		speaker.Play(a.stream())
	}
}

// Rebuilds the chain of controls and effects.
// rawDataStream -> Loop (optional) -> Resample (optional) -> Ctrl -> Volume -> Pan
func (a *Audible) refreshEffectsChain() {
	var looped beep.Streamer
	if a.Looping() {
		// -1 loops forever
		looped = beep.Loop(-1, a.rawDataStream)
	} else {
		// no looping requested
		looped = a.rawDataStream
	}

	var resampled beep.Streamer
	if a.rawDataFormat.SampleRate != SampleRate {
		resampled = beep.Resample(4, a.rawDataFormat.SampleRate, SampleRate, looped)
	} else {
		// if the sample rate of the raw audio data matches our speaker setup, no resampling necessary
		resampled = looped
	}

	a.controller.Streamer = resampled
}

// stream returns the top-level Streamer at the end of the controls and effects chain.
// Additionally, wraps the whole chain in a callback that alerts us when the Audible has finished playing,
// and runs the FinishedCallback if it is set.
func (a *Audible) stream() beep.Streamer {
	return beep.Seq(a.topLevelStream, beep.Callback(func() {
		a.Pause()
		a.Rewind()
		a.active = false

		if a.finishedCallback != nil {
			a.finishedCallback(a)
		}
	}))
}

// Active returns true if the Audible is currently Active,
// meaning that the Audible's Play method has been called, and it has not yet finished playing.
// If the Audible is not Active, then pausing/unpausing, panning, looping, etc., will appear to have no effect
// until the Audible's Play method is called.
// When an Audible is Played, it is marked Active, and then it is automatically marked inactive when it has finished
// playing.
func (a *Audible) Active() bool {
	return a.active
}

// Play unpauses the Audible if it is paused, and re-activates it if it is inactive.
// The Audible is automatically marked inactive and rewinded when it finishes playing.
// If Play is called again before the Audible has finished playing, this Audible will appear to stop and restart back
// at the beginning of the stream.
func (a *Audible) Play() {
	a.Unpause()

	if !a.active {
		a.active = true
		speaker.Play(a.stream())
	}
}

// Rewind rewinds the Audible's data stream back to the beginning.
func (a *Audible) Rewind() {
	if a.rawDataStream != nil {
		a.rawDataStream.Seek(0)
	}
}

// Paused returns true if the Audible is paused.
func (a *Audible) Paused() bool {
	return a.controller.Paused
}

// Pause pauses the Audible.
func (a *Audible) Pause() {
	a.controller.Paused = true
}

// Unpause unpauses the Audible.
func (a *Audible) Unpause() {
	a.controller.Paused = false
}

// Volume returns the Audible's current volume. This will return the same value regardless of whether the Audible
// is muted.
func (a *Audible) Volume() float64 {
	return a.volumeController.Volume
}

// SetVolume sets the Audible's volume. It is recommended that you start at 0 and go up or down in increments of 0.5.
func (a *Audible) SetVolume(volume float64) {
	a.volumeController.Volume = volume
}

// Muted returns true if the Audible is muted.
func (a *Audible) Muted() bool {
	return a.volumeController.Silent
}

// Mute mutes the Audible, making it completely silent.
func (a *Audible) Mute() {
	a.volumeController.Silent = true
}

// Unmute unmutes the Audible.
func (a *Audible) Unmute() {
	a.volumeController.Silent = false
}

// Looping returns true if the Audible is set to loop endlessly.
func (a *Audible) Looping() bool {
	return a.looping
}

// SetLooping takes a boolean argument, and sets the Audible to loop endlessly if true.
func (a *Audible) SetLooping(loop bool) {
	// if we weren't looping and now we are, or vice versa, we need to alter the effects chain to add/remove the looper
	if loop != a.Looping() {
		a.looping = loop

		a.refreshEffectsChain()
	}
}

// Pan returns a float between -1 and 1 representing the current pan of the audio.
// -1 = 100% left channel
// 0 = 50% left/50% right
// 1 = 100% right channel
func (a *Audible) Pan() float64 {
	return a.panController.Pan
}

// SetPan takes a float between -1 and 1 to control the left-right pan of the audio.
// -1 = 100% left channel
// 0 = 50% left/50% right
// 1 = 100% right channel
func (a *Audible) SetPan(pan float64) {
	if pan < -1 {
		pan = -1
	} else if pan > 1 {
		pan = 1
	}

	a.panController.Pan = pan
}

// SetFinishedCallback sets the callback function that is called when the Audible has finished playing.
// This callback function is never called if the Audible is looping.
func (a *Audible) SetFinishedCallback(callbackFn func(*Audible)) {
	a.finishedCallback = callbackFn
}

// Progress returns a float between 0 and 100 representing the percentage of the Audible that has finished playing.
func (a *Audible) Progress() float64 {
	return float64(a.rawDataStream.Position()) / float64(a.rawDataStream.Len()) * 100
}

// Duration returns the length of the Audible's raw audio data as a time.Duration
func (a *Audible) Duration() time.Duration {
	return a.rawDataFormat.SampleRate.D(a.rawDataStream.Len())
}

// Position returns Audible's current position in its raw audio data stream as a time.Duration
func (a *Audible) Position() time.Duration {
	return a.rawDataFormat.SampleRate.D(a.rawDataStream.Position())
}

// AudibleFactory is a wrapper for the generic component factory that returns Audible component instances.
// This can be embedded inside a system to give them the methods for adding, retrieving, and removing an Audible.
type AudibleFactory struct {
	*akara.ComponentFactory
}

// Add adds an Audible component to the given entity and returns it
func (m *AudibleFactory) Add(id akara.EID) *Audible {
	return m.ComponentFactory.Add(id).(*Audible)
}

// Get returns the Audible component for the given entity, and a bool for whether not it exists
func (m *AudibleFactory) Get(id akara.EID) (*Audible, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Audible), found
}
