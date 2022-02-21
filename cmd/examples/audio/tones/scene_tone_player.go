package main

import (
	"fmt"
	"image"
	"image/color"

	. "github.com/gravestench/director"

	"github.com/gravestench/director/pkg/components/audio"
	"github.com/gravestench/director/pkg/components/text"

	"github.com/gravestench/director/pkg/systems/input/constants"
	"github.com/gravestench/director/pkg/systems/input/vector"
)

const (
	key = "Director Example - Audio Tone Player"
)

type AudioTonePlayerScene struct {
	Scene
	sounds         []*audio.Audio
	infoTextLabels []*text.Component
	loopingText    *text.Component
}

func (scene *AudioTonePlayerScene) Key() string {
	return key
}

func (scene *AudioTonePlayerScene) Init(_ *World) {
	scene.loadSounds()
	scene.createPlayButtons()
	scene.createLoopButton()
	scene.createVolumeButtons()
	scene.createPanButtons()
	scene.createSpeedButtons()

	scene.Sys.Renderer.Window.Title = scene.Key()
}

func (scene *AudioTonePlayerScene) loadSounds() {
	eid1 := scene.Add.Sound("data/220hz.wav", true, 0, false, false, 1)
	eid2 := scene.Add.Sound("data/440hz.wav", true, 0, false, false, 1)

	audible1, _ := scene.Components.Audible.Get(eid1)
	audible2, _ := scene.Components.Audible.Get(eid2)
	scene.sounds = append(scene.sounds, audible1, audible2)
}

func (scene *AudioTonePlayerScene) createPlayButtons() {
	red := color.RGBA{R: 255, A: 255}
	blue := color.RGBA{B: 255, A: 255}
	bWidth, bHeight := 200, 100

	lowFreqButtonEid := scene.Add.Rectangle(150, 100, bWidth, bHeight, red, red)
	highFreqButtonEid := scene.Add.Rectangle(450, 100, bWidth, bHeight, blue, blue)

	scene.Add.Label("220Hz", 150, 120, 24, "", color.Black)
	scene.Add.Label("440Hz", 450, 120, 24, "", color.Black)

	lowFreqButtonInfoTextEid := scene.Add.Label("", 190, 70, 18, "", color.Black)
	highFreqButtonInfoTextEid := scene.Add.Label("", 490, 70, 18, "", color.Black)
	lowFreqButtonInfoText, _ := scene.Components.Text.Get(lowFreqButtonInfoTextEid)
	highFreqButtonInfoText, _ := scene.Components.Text.Get(highFreqButtonInfoTextEid)
	scene.infoTextLabels = append(scene.infoTextLabels, lowFreqButtonInfoText, highFreqButtonInfoText)

	lowFreqButtonInteractive := scene.Components.Interactive.Add(lowFreqButtonEid)
	highFreqButtonInteractive := scene.Components.Interactive.Add(highFreqButtonEid)

	lowFreqButtonInteractive.Vector = vector.NewInputVector()
	lowFreqButtonInteractive.Vector.SetMouseButton(constants.MouseButtonLeft)
	lowFreqButtonInteractive.Callback = func() (preventPropogation bool) {
		if scene.sounds[0].Paused() {
			scene.sounds[0].Play()
		} else {
			scene.sounds[0].Pause()
		}
		return false
	}
	lowFreqButtonInteractive.Hitbox = scene.getHitboxRectangle(lowFreqButtonEid)
	lowFreqButtonInteractive.Enabled = true

	highFreqButtonInteractive.Vector = vector.NewInputVector()
	highFreqButtonInteractive.Vector.SetMouseButton(constants.MouseButtonLeft)
	highFreqButtonInteractive.Callback = func() (preventPropogation bool) {
		if scene.sounds[1].Paused() {
			scene.sounds[1].Play()
		} else {
			scene.sounds[1].Pause()
		}
		return false
	}
	highFreqButtonInteractive.Hitbox = scene.getHitboxRectangle(highFreqButtonEid)
	highFreqButtonInteractive.Enabled = true

}

func (scene *AudioTonePlayerScene) createLoopButton() {
	loopingTextEid := scene.Add.Label("Loop (off)", 350, 350, 24, "", color.White)
	scene.loopingText, _ = scene.Components.Text.Get(loopingTextEid)

	loopButtonEid := scene.Add.Rectangle(350, 300, 50, 33, color.RGBA{R: 0, G: 180, B: 0, A: 255}, color.RGBA{R: 0, G: 180, B: 0, A: 255})

	scene.Add.Label("<-", 750, 100, 24, "", color.White)

	loopButtonInteractive := scene.Components.Interactive.Add(loopButtonEid)

	loopButtonInteractive.Vector = vector.NewInputVector()
	loopButtonInteractive.Vector.SetMouseButton(constants.MouseButtonLeft)
	loopButtonInteractive.Callback = func() (preventPropogation bool) {
		for _, sound := range scene.sounds {
			sound.SetLooping(!sound.Looping())
		}
		return false
	}
	loopButtonInteractive.Hitbox = scene.getHitboxRectangle(loopButtonEid)
	loopButtonInteractive.Enabled = true
}

func (scene *AudioTonePlayerScene) createPanButtons() {
	scene.Add.Label("Pan", 800, 130, 24, "", color.White)
	panLeftButtonEid := scene.Add.Rectangle(750, 100, 50, 33, color.RGBA{R: 0, G: 180, B: 0, A: 255}, color.RGBA{R: 0, G: 180, B: 0, A: 255})
	panRightButtonEid := scene.Add.Rectangle(825, 100, 50, 33, color.RGBA{R: 0, G: 180, B: 0, A: 255}, color.RGBA{R: 0, G: 180, B: 0, A: 255})

	scene.Add.Label("<-", 750, 100, 24, "", color.White)
	scene.Add.Label("->", 825, 100, 24, "", color.White)

	panLeftButtonInteractive := scene.Components.Interactive.Add(panLeftButtonEid)
	panRightButtonInteractive := scene.Components.Interactive.Add(panRightButtonEid)

	panLeftButtonInteractive.Vector = vector.NewInputVector()
	panLeftButtonInteractive.Vector.SetMouseButton(constants.MouseButtonLeft)
	panLeftButtonInteractive.Callback = func() (preventPropogation bool) {
		for _, sound := range scene.sounds {
			sound.SetPan(sound.Pan() - 0.1)
		}
		return false
	}
	panLeftButtonInteractive.Hitbox = scene.getHitboxRectangle(panLeftButtonEid)
	panLeftButtonInteractive.Enabled = true

	panRightButtonInteractive.Vector = vector.NewInputVector()
	panRightButtonInteractive.Vector.SetMouseButton(constants.MouseButtonLeft)
	panRightButtonInteractive.Callback = func() (preventPropogation bool) {
		for _, sound := range scene.sounds {
			sound.SetPan(sound.Pan() + 0.1)
		}
		return false
	}
	panRightButtonInteractive.Hitbox = scene.getHitboxRectangle(panRightButtonEid)
	panRightButtonInteractive.Enabled = true
}

func (scene *AudioTonePlayerScene) createVolumeButtons() {
	scene.Add.Label("Volume", 800, 430, 24, "", color.White)
	volumeUpButtonEid := scene.Add.Rectangle(750, 400, 50, 33, color.RGBA{R: 0, G: 180, B: 0, A: 255}, color.RGBA{R: 0, G: 180, B: 0, A: 255})
	volumeDownButtonEid := scene.Add.Rectangle(825, 400, 50, 33, color.RGBA{R: 0, G: 180, B: 0, A: 255}, color.RGBA{R: 0, G: 180, B: 0, A: 255})

	scene.Add.Label("^", 750, 400, 24, "", color.White)
	scene.Add.Label("v", 825, 400, 24, "", color.White)

	volumeUpButtonInteractive := scene.Components.Interactive.Add(volumeUpButtonEid)
	volumeDownButtonInteractive := scene.Components.Interactive.Add(volumeDownButtonEid)

	volumeUpButtonInteractive.Vector = vector.NewInputVector()
	volumeUpButtonInteractive.Vector.SetMouseButton(constants.MouseButtonLeft)
	volumeUpButtonInteractive.Callback = func() (preventPropogation bool) {
		for _, sound := range scene.sounds {
			sound.SetVolume(sound.Volume() + 0.5)
		}
		return false
	}
	volumeUpButtonInteractive.Hitbox = scene.getHitboxRectangle(volumeUpButtonEid)
	volumeUpButtonInteractive.Enabled = true

	volumeDownButtonInteractive.Vector = vector.NewInputVector()
	volumeDownButtonInteractive.Vector.SetMouseButton(constants.MouseButtonLeft)
	volumeDownButtonInteractive.Callback = func() (preventPropogation bool) {
		for _, sound := range scene.sounds {
			sound.SetVolume(sound.Volume() - 0.5)
		}
		return false
	}
	volumeDownButtonInteractive.Hitbox = scene.getHitboxRectangle(volumeDownButtonEid)
	volumeDownButtonInteractive.Enabled = true
}

func (scene *AudioTonePlayerScene) createSpeedButtons() {
	scene.Add.Label("Speed", 800, 730, 24, "", color.White)
	speedUpButtonEid := scene.Add.Rectangle(750, 700, 50, 33, color.RGBA{R: 0, G: 180, B: 0, A: 255}, color.RGBA{R: 0, G: 180, B: 0, A: 255})
	speedDownButtonEid := scene.Add.Rectangle(825, 700, 50, 33, color.RGBA{R: 0, G: 180, B: 0, A: 255}, color.RGBA{R: 0, G: 180, B: 0, A: 255})

	scene.Add.Label("^", 750, 700, 24, "", color.White)
	scene.Add.Label("v", 825, 700, 24, "", color.White)

	speedUpButtonInteractive := scene.Components.Interactive.Add(speedUpButtonEid)
	speedDownButtonInteractive := scene.Components.Interactive.Add(speedDownButtonEid)

	speedUpButtonInteractive.Vector = vector.NewInputVector()
	speedUpButtonInteractive.Vector.SetMouseButton(constants.MouseButtonLeft)
	speedUpButtonInteractive.Callback = func() (preventPropogation bool) {
		for _, sound := range scene.sounds {
			sound.SetSpeedMultiplier(sound.SpeedMultiplier() + 0.1)
		}
		return false
	}
	speedUpButtonInteractive.Hitbox = scene.getHitboxRectangle(speedUpButtonEid)
	speedUpButtonInteractive.Enabled = true

	speedDownButtonInteractive.Vector = vector.NewInputVector()
	speedDownButtonInteractive.Vector.SetMouseButton(constants.MouseButtonLeft)
	speedDownButtonInteractive.Callback = func() (preventPropogation bool) {
		for _, sound := range scene.sounds {
			sound.SetSpeedMultiplier(sound.SpeedMultiplier() - 0.1)
		}
		return false
	}
	speedDownButtonInteractive.Hitbox = scene.getHitboxRectangle(speedDownButtonEid)
	speedDownButtonInteractive.Enabled = true
}

func (scene *AudioTonePlayerScene) updateInfoText() {
	for i, sound := range scene.sounds {
		playingOrPaused := "PLAYING"
		if sound.Paused() || !sound.Active() {
			playingOrPaused = "PAUSED"
		}

		scene.infoTextLabels[i].String = fmt.Sprintf("%s\n%.1fs / %.1fs (%.0f%%)",
			playingOrPaused, sound.Position().Seconds(), sound.Duration().Seconds(), sound.Progress())
	}
}

func (scene *AudioTonePlayerScene) updateLoopingText() {
	looping := "off"
	if scene.sounds[0].Looping() {
		looping = "on"
	}

	scene.loopingText.String = fmt.Sprintf("Loop (%s)", looping)
}

func (scene *AudioTonePlayerScene) Update() {
	scene.updateInfoText()
	scene.updateLoopingText()
}

// we probably need a better way to do this
func (scene *AudioTonePlayerScene) getHitboxRectangle(id Entity) *image.Rectangle {
	size, _ := scene.Components.Size.Get(id)
	trs, _ := scene.Components.Transform.Get(id)
	rHeight := scene.Sys.Renderer.Window.Height

	return &image.Rectangle{
		Min: image.Point{
			X: int(trs.Translation.X) - size.Dx()/2,
			Y: rHeight - (int(trs.Translation.Y) + size.Dy()/2),
		},
		Max: image.Point{
			X: int(trs.Translation.X) + size.Dx()/2,
			Y: rHeight - (int(trs.Translation.Y) - size.Dy()/2),
		},
	}
}
