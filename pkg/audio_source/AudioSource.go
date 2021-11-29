package audio_source

import (
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/speaker"
)

const (
	SampleRate        = 48000
	BufferSizeSamples = 2400
)

func New() *AudioSource {
	audio := &AudioSource{}

	audio.controller = &beep.Ctrl{Streamer: audio.rawDataStream}

	audio.volumeController = &effects.Volume{
		Streamer: audio.controller,
		Base:     2,
	}

	audio.panController = &effects.Pan{
		Streamer: audio.volumeController,
	}

	audio.speedController = beep.ResampleRatio(4, 1, audio.panController)

	audio.topLevelStream = audio.speedController

	return audio
}

// AudioSource expresses the state of audio playback for a single source of audio data.
type AudioSource struct {
	// active denotes whether the speaker holds a reference to this Audio
	active           bool
	looping          bool
	rawDataStream    beep.StreamSeekCloser
	rawDataFormat    beep.Format
	controller       *beep.Ctrl
	volumeController *effects.Volume
	panController    *effects.Pan
	speedController  *beep.Resampler
	topLevelStream   beep.Streamer

	// FinishedCallback is called when the Audio has finished playing. This is never called if the Audio is looping.
	// Receives one argument, the Audio itself.
	finishedCallback func(*AudioSource)
}

func (a *AudioSource) SetStream(stream beep.StreamSeekCloser, format beep.Format) {
	a.rawDataStream = stream
	a.rawDataFormat = format

	a.refreshEffectsChain()

	// immediately play the Audio if it was not created in the Paused state
	if !a.Paused() {
		speaker.Play(a.stream())
	}
}

// Rebuilds the chain of controls and effects.
// rawDataStream -> Loop (optional) -> Resample (optional) -> Ctrl -> Volume -> Pan -> Speed
func (a *AudioSource) refreshEffectsChain() {
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
// Additionally, wraps the whole chain in a callback that alerts us when the Audio has finished playing,
// and runs the FinishedCallback if it is set.
func (a *AudioSource) stream() beep.Streamer {
	return beep.Seq(a.topLevelStream, beep.Callback(func() {
		a.Pause()
		a.Rewind()
		a.active = false

		if a.finishedCallback != nil {
			a.finishedCallback(a)
		}
	}))
}

// Active returns true if the Audio is currently Active,
// meaning that the Audio's Play method has been called, and it has not yet finished playing.
// If the Audio is not Active, then pausing/unpausing, panning, looping, etc., will appear to have no effect
// until the Audio's Play method is called.
// When an Audio is Played, it is marked Active, and then it is automatically marked inactive when it has finished
// playing.
func (a *AudioSource) Active() bool {
	return a.active
}

// Play unpauses the Audio if it is paused, and re-activates it if it is inactive.
// The Audio is automatically marked inactive and rewinded when it finishes playing.
// If Play is called again before the Audio has finished playing, this Audio will appear to stop and restart back
// at the beginning of the stream.
func (a *AudioSource) Play() {
	a.Unpause()

	if !a.active {
		a.active = true
		speaker.Play(a.stream())
	}
}

// Rewind rewinds the Audio's data stream back to the beginning.
func (a *AudioSource) Rewind() {
	if a.rawDataStream != nil {
		a.rawDataStream.Seek(0)
	}
}

// Paused returns true if the Audio is paused.
func (a *AudioSource) Paused() bool {
	return a.controller.Paused
}

// Pause pauses the Audio.
func (a *AudioSource) Pause() {
	a.controller.Paused = true
}

// Unpause unpauses the Audio.
func (a *AudioSource) Unpause() {
	a.controller.Paused = false
}

// Volume returns the Audio's current volume. This will return the same value regardless of whether the Audio
// is muted.
func (a *AudioSource) Volume() float64 {
	return a.volumeController.Volume
}

// SetVolume sets the Audio's volume. It is recommended that you start at 0 and go up or down in increments of 0.5.
func (a *AudioSource) SetVolume(volume float64) {
	a.volumeController.Volume = volume
}

// Muted returns true if the Audio is muted.
func (a *AudioSource) Muted() bool {
	return a.volumeController.Silent
}

// Mute mutes the Audio, making it completely silent.
func (a *AudioSource) Mute() {
	a.volumeController.Silent = true
}

// Unmute unmutes the Audio.
func (a *AudioSource) Unmute() {
	a.volumeController.Silent = false
}

// Looping returns true if the Audio is set to loop endlessly.
func (a *AudioSource) Looping() bool {
	return a.looping
}

// SetLooping takes a boolean argument, and sets the Audio to loop endlessly if true.
func (a *AudioSource) SetLooping(loop bool) {
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
func (a *AudioSource) Pan() float64 {
	return a.panController.Pan
}

// SetPan takes a float between -1 and 1 to control the left-right pan of the audio.
// -1 = 100% left channel
// 0 = 50% left/50% right
// 1 = 100% right channel
func (a *AudioSource) SetPan(pan float64) {
	if pan < -1 {
		pan = -1
	} else if pan > 1 {
		pan = 1
	}

	a.panController.Pan = pan
}

// SpeedMultiplier returns the Audio's current speed multiplier.
func (a *AudioSource) SpeedMultiplier() float64 {
	return a.speedController.Ratio()
}

// SetSpeedMultiplier sets the Audio's speed multiplier.
// Examples:
// 1 = normal speed
// 2 = 2x normal speed
// 0.5 = 0.5x normal speed
func (a *AudioSource) SetSpeedMultiplier(multiplier float64) {
	if multiplier < 0 {
		multiplier = 0.01
	}
	a.speedController.SetRatio(multiplier)
}

// SetFinishedCallback sets the callback function that is called when the Audio has finished playing.
// This callback function is never called if the Audio is looping.
func (a *AudioSource) SetFinishedCallback(callbackFn func(*AudioSource)) {
	a.finishedCallback = callbackFn
}

// Progress returns a float between 0 and 100 representing the percentage of the Audio that has finished playing.
func (a *AudioSource) Progress() float64 {
	return float64(a.rawDataStream.Position()) / float64(a.rawDataStream.Len()) * 100
}

// Duration returns the length of the Audio's raw audio data as a time.Duration
func (a *AudioSource) Duration() time.Duration {
	return a.rawDataFormat.SampleRate.D(a.rawDataStream.Len())
}

// Position returns Audio's current position in its raw audio data stream as a time.Duration
func (a *AudioSource) Position() time.Duration {
	return a.rawDataFormat.SampleRate.D(a.rawDataStream.Position())
}
