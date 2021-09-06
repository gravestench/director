package main

import (
	"fmt"
	"image"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/systems/input"

	"github.com/gravestench/director/pkg/systems/scene"
)

type GameScene struct {
	scene.Scene
	toggleButton   akara.EID
	toggleLabel    akara.EID
	debugPanel     akara.EID
	mainPanel      akara.EID
	square         akara.EID
	label          akara.EID
	testLabel      akara.EID
	isDebugEnabled bool
}

// Game Loop
func (scene *GameScene) Update() {
	scene.updateTestLabel()
	//scene.updateLabel()

}

func (scene *GameScene) makeMainPanel() {
	background := color.RGBA{R: 30, G: 31, B: 35, A: 255}
	scene.toggleButton = scene.Add.Rectangle(scene.Window.Width/2, scene.Window.Height/2, scene.Window.Width, scene.Window.Height, background, nil)
}

/****************************
*	  Toggle debug code		*
****************************/

func (scene *GameScene) makeToggleButton() {
	purple := color.RGBA{R: 104, G: 70, B: 236, A: 255}
	scene.toggleButton = scene.Add.Rectangle(scene.Window.Width-60, scene.Window.Height-15, 140, 30, purple, nil)
}

func (scene *GameScene) toggleDebug() {
	if scene.isDebugEnabled == false {
		scene.isDebugEnabled = true
		scene.clearDebugPanel()
		scene.makeDebugPanel()
		fmt.Println(scene.debugPanel)
	} else {
		scene.isDebugEnabled = false
		scene.clearDebugPanel()
		fmt.Println(scene.debugPanel)
	}
}

func (scene *GameScene) clearDebugPanel() {
	fmt.Print("Remove Debug Panel")
	scene.Director.RemoveEntity(scene.debugPanel)
}

func (scene *GameScene) makeDebugPanel() {
	scene.Director.RemoveEntity(scene.debugPanel)
	background := color.RGBA{R: 21, G: 23, B: 24, A: 255}
	scene.debugPanel = scene.Add.Rectangle(scene.Window.Width/4, scene.Window.Height/4, scene.Window.Width, scene.Window.Height, background, nil)
}

func (scene *GameScene) makeToggleLabel() {
	tgBtnTrs, found := scene.Components.Transform.Get(scene.toggleButton)
	if !found {
		return
	}
	fmt.Print(tgBtnTrs)

	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	scene.toggleLabel = scene.Add.Label("Toggle Debug", scene.Window.Width-30, scene.Window.Height-15, 12, "", white)
}

func (scene *GameScene) bindDebugInput() {
	i := scene.Components.Interactive.Add(scene.toggleButton)

	i.Callback = func() (preventPropogation bool) {
		scene.toggleDebug()
		fmt.Print("Test")
		return false
	}

	i.Vector = input.NewInputVector()
	i.Vector.SetMouseButton(input.MouseButtonLeft)

	size, found := scene.Components.Size.Get(scene.toggleButton)
	if !found {
		return
	}

	trs, found := scene.Components.Transform.Get(scene.toggleButton)
	if !found {
		return
	}

	i.CursorPosition = &image.Rectangle{
		Min: image.Point{
			X: int(trs.Translation.X) - size.Dx()/2,
			Y: scene.Window.Height - (int(trs.Translation.Y) + size.Dy()/2),
		},
		Max: image.Point{
			X: int(trs.Translation.X) + size.Dx()/2,
			Y: scene.Window.Height - (int(trs.Translation.Y) - size.Dy()/2),
		},
	}

	fmt.Print(i)
}

/****************************
*	End toggle debug code	*
****************************/

func (scene *GameScene) updateTestLabel() {
	text, found := scene.Components.Text.Get(scene.testLabel)
	if !found {
		return
	}

	mp := rl.GetMousePosition()

	const (
		fmtMouse = "Mouse (%v, %v)"
	)

	text.String = fmt.Sprintf(fmtMouse, mp.X, mp.Y)
}

func (scene *GameScene) updateLabel() {
	text, found := scene.Components.Text.Get(scene.label)
	if !found {
		return
	}

	mp := rl.GetMousePosition()

	const (
		fmtMouse = "Mouse (%v, %v)"
	)

	text.String = fmt.Sprintf(fmtMouse, mp.X, mp.Y)
}

func (scene *GameScene) Init(world *akara.World) {
	scene.isDebugEnabled = false
	scene.makeMainPanel()
	scene.makeToggleButton()
	scene.makeToggleLabel()
	scene.makeLabel()

	scene.bindDebugInput()
}

func (scene *GameScene) makeSquare() {
	blue := color.RGBA{B: 255, A: 255}
	scene.square = scene.Add.Rectangle(100, 100, 30, 30, blue, nil)
}

func (scene *GameScene) makeLabel() {
	red := color.RGBA{R: 255, A: 255}
	scene.label = scene.Add.Label("", 400, 400, 24, "", red)
	scene.testLabel = scene.Add.Label("", 200, 200, 24, "", red)
}

func (scene *GameScene) IsInitialized() bool {

	return scene.toggleButton != 0
}
