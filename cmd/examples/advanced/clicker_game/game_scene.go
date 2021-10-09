package main

import (
	"fmt"
	"github.com/gravestench/director/pkg/common"
	"image"
	"image/color"
	"strconv"

	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/systems/input"

	"github.com/gravestench/director/pkg/systems/scene"
)

type GameScene struct {
	scene.Scene
	upgrades     ShopUpgrades
	toggleButton common.Entity
	toggleLabel  common.Entity
	debugPanel   common.Entity
	//mainPanel       common.Entity
	//square          common.Entity
	//label           common.Entity
	mouseDebugLabel common.Entity
	balanceLabel    common.Entity
	clickButton     common.Entity
	shopPanel       common.Entity
	isDebugEnabled  bool
	balanceValue    int
	clickValue      int
}

type ShopUpgrades struct {
	clickerUpgrade1      common.Entity
	clickerUpgrade1Label common.Entity
	clickerUpgrade2      common.Entity
	clickerUpgrade2Label common.Entity
	clickerUpgrade3      common.Entity
	clickerUpgrade3Label common.Entity
	clickerUpgrade4      common.Entity
	clickerUpgrade4Label common.Entity
	//clickerUpgrade1Price int
	//clickerUpgrade2Price int
	//clickerUpgrade3Price int
	//clickerUpgrade4Price int
}

func (scene *GameScene) Key() string {
	return "Game Test"
}

// Game Loop
func (scene *GameScene) Update() {
	//scene.updateLabel()
	if scene.isDebugEnabled {
		scene.updateTestLabel()
	}
}

func (scene *GameScene) makeMainPanel() {
	background := color.RGBA{R: 30, G: 31, B: 35, A: 255}
	rWidth := scene.Sys.Renderer.Window.Width
	rHeight := scene.Sys.Renderer.Window.Height

	scene.toggleButton = scene.Add.Rectangle(rWidth/2, rHeight/2, rWidth, rHeight, background, nil)
}

/****************************
*	  Toggle debug code		*
****************************/
func (scene *GameScene) makeToggleButton() {
	purple := color.RGBA{R: 104, G: 70, B: 236, A: 255}
	rWidth := scene.Sys.Renderer.Window.Width
	rHeight := scene.Sys.Renderer.Window.Height

	scene.toggleButton = scene.Add.Rectangle(rWidth-60, rHeight-15, 140, 30, purple, nil)
}
func (scene *GameScene) toggleDebug() {
	if !scene.isDebugEnabled {
		scene.isDebugEnabled = true
		scene.makeDebugPanel()
		scene.makeMouseDebugLabel()
	} else {
		scene.isDebugEnabled = false
		scene.clearDebugPanel()
	}
}
func (scene *GameScene) clearDebugPanel() {
	scene.RemoveEntity(scene.debugPanel)
	scene.RemoveEntity(scene.mouseDebugLabel)
}
func (scene *GameScene) makeDebugPanel() {
	background := color.RGBA{R: 21, G: 23, B: 24, A: 255}
	rWidth := scene.Sys.Renderer.Window.Width
	rHeight := scene.Sys.Renderer.Window.Height

	scene.debugPanel = scene.Add.Rectangle(rWidth/2-132, rHeight-30, rWidth, 60, background, nil)
}
func (scene *GameScene) makeToggleLabel() {
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	rWidth := scene.Sys.Renderer.Window.Width
	rHeight := scene.Sys.Renderer.Window.Height

	scene.toggleLabel = scene.Add.Label("Toggle Debug", rWidth-30, rHeight-15, 12, "", white)
}
func (scene *GameScene) makeMouseDebugLabel() {
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	rWidth := scene.Sys.Renderer.Window.Width
	rHeight := scene.Sys.Renderer.Window.Height

	scene.mouseDebugLabel = scene.Add.Label("Mouse: ", rWidth-30, rHeight-15, 12, "", white)
	origin, found := scene.Components.Origin.Get(scene.mouseDebugLabel)
	if !found {
		return
	}
	origin.X = 0
	origin.Y = 0
}

func (scene *GameScene) bindDebugInput() {
	i := scene.Components.Interactive.Add(scene.toggleButton)

	i.Callback = func() (preventPropagation bool) {
		scene.toggleDebug()
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

	rHeight := scene.Sys.Renderer.Window.Height

	i.Hitbox = &image.Rectangle{
		Min: image.Point{
			X: int(trs.Translation.X) - size.Dx()/2,
			Y: rHeight - (int(trs.Translation.Y) + size.Dy()/2),
		},
		Max: image.Point{
			X: int(trs.Translation.X) + size.Dx()/2,
			Y: rHeight - (int(trs.Translation.Y) - size.Dy()/2),
		},
	}

	fmt.Print(i)
}
func (scene *GameScene) updateTestLabel() {
	text, found := scene.Components.Text.Get(scene.mouseDebugLabel)
	if !found {
		return
	}

	const (
		fmtMouse = "Mouse: (%v, %v)"
	)

	text.String = fmt.Sprintf(fmtMouse, scene.Sys.Input.MousePosition.X, scene.Sys.Input.MousePosition.Y)
}

/****************************
*	End toggle debug code	*
****************************/

func (scene *GameScene) makeInitialUI() {
	gray := color.RGBA{R: 97, G: 98, B: 109, A: 255}

	rWidth := scene.Sys.Renderer.Window.Width
	rHeight := scene.Sys.Renderer.Window.Height

	scene.makeMainPanel()
	scene.makeToggleButton()
	scene.makeToggleLabel()
	scene.makeShopPanel()
	scene.balanceLabel = scene.Add.Label("Balance: 0 Cubes", rWidth/2+40, rHeight/2+85, 24, "", gray)
	scene.makeClickButton()
}

func (scene *GameScene) makeClickButton() {
	purple := color.RGBA{R: 104, G: 70, B: 236, A: 255}
	rWidth := scene.Sys.Renderer.Window.Width
	rHeight := scene.Sys.Renderer.Window.Height

	scene.clickButton = scene.Add.Rectangle(rWidth/2, rHeight/2, 140, 140, purple, nil)
}

func (scene *GameScene) makeShopPanel() {
	dark := color.RGBA{R: 13, G: 12, B: 15, A: 255}
	shopWidth := 150
	rHeight := scene.Sys.Renderer.Window.Height

	scene.shopPanel = scene.Add.Rectangle(shopWidth/2, rHeight/2, shopWidth, rHeight, dark, nil)
	scene.makeShopUpgrades()
}

func (scene *GameScene) makeShopUpgrades() {
	rHeight := scene.Sys.Renderer.Window.Height
	upgradeYLocation := rHeight - 20
	upgradeSize := 30
	shopWidth := 150
	purple := color.RGBA{R: 104, G: 70, B: 236, A: 255}
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	scene.upgrades.clickerUpgrade1 = scene.Add.Rectangle(shopWidth/2, upgradeYLocation, shopWidth-10, upgradeSize-5, purple, nil)
	scene.upgrades.clickerUpgrade1Label = scene.Add.Label("Upgrade 1", shopWidth/2, upgradeYLocation, 12, "", white)
	upgradeYLocation -= upgradeSize
	scene.upgrades.clickerUpgrade2 = scene.Add.Rectangle(shopWidth/2, upgradeYLocation, shopWidth-10, upgradeSize-5, purple, nil)
	scene.upgrades.clickerUpgrade2Label = scene.Add.Label("Upgrade 2", shopWidth/2, upgradeYLocation, 12, "", white)
	upgradeYLocation -= upgradeSize
	scene.upgrades.clickerUpgrade3 = scene.Add.Rectangle(shopWidth/2, upgradeYLocation, shopWidth-10, upgradeSize-5, purple, nil)
	scene.upgrades.clickerUpgrade3Label = scene.Add.Label("Upgrade 3", shopWidth/2, upgradeYLocation, 12, "", white)
	upgradeYLocation -= upgradeSize
	scene.upgrades.clickerUpgrade4 = scene.Add.Rectangle(shopWidth/2, upgradeYLocation, shopWidth-10, upgradeSize-5, purple, nil)
	scene.upgrades.clickerUpgrade4Label = scene.Add.Label("Upgrade 4", shopWidth/2, upgradeYLocation, 12, "", white)
}

func (scene *GameScene) upgradeClicker(value int) {
	scene.clickValue += value
}

func (scene *GameScene) bindClickingInput() {
	i := scene.Components.Interactive.Add(scene.clickButton)
	i.Callback = func() (preventPropagation bool) {
		scene.updateBalance(scene.clickValue)
		return false
	}
	i.Vector = input.NewInputVector()
	i.Vector.SetMouseButton(input.MouseButtonLeft)
	size, _ := scene.Components.Size.Get(scene.clickButton)
	trs, _ := scene.Components.Transform.Get(scene.clickButton)
	rHeight := scene.Sys.Renderer.Window.Height

	i.Hitbox = &image.Rectangle{
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

func (scene *GameScene) bindShopClickingInput() {
	i := scene.Components.Interactive.Add(scene.upgrades.clickerUpgrade1)
	i.Callback = func() (preventPropagation bool) {
		fmt.Print(scene.clickValue)
		scene.upgradeClicker(1)
		scene.RemoveEntity(scene.upgrades.clickerUpgrade1)
		scene.RemoveEntity(scene.upgrades.clickerUpgrade1Label)
		return false
	}
	i.Vector = input.NewInputVector()
	i.Vector.SetMouseButton(input.MouseButtonLeft)
	size, found := scene.Components.Size.Get(scene.upgrades.clickerUpgrade1)
	if !found {
		return
	}

	trs, found := scene.Components.Transform.Get(scene.upgrades.clickerUpgrade1)
	if !found {
		return
	}

	rHeight := scene.Sys.Renderer.Window.Height

	i.Hitbox = &image.Rectangle{
		Min: image.Point{
			X: int(trs.Translation.X) - size.Dx()/2,
			Y: rHeight - (int(trs.Translation.Y) + size.Dy()/2),
		},
		Max: image.Point{
			X: int(trs.Translation.X) + size.Dx()/2,
			Y: rHeight - (int(trs.Translation.Y) - size.Dy()/2),
		},
	}

	i = scene.Components.Interactive.Add(scene.upgrades.clickerUpgrade2)
	i.Callback = func() (preventPropagation bool) {
		scene.upgradeClicker(2)
		scene.RemoveEntity(scene.upgrades.clickerUpgrade2)
		scene.RemoveEntity(scene.upgrades.clickerUpgrade2Label)
		return false
	}
	i.Vector = input.NewInputVector()
	i.Vector.SetMouseButton(input.MouseButtonLeft)
	size, found = scene.Components.Size.Get(scene.upgrades.clickerUpgrade2)
	if !found {
		return
	}

	trs, found = scene.Components.Transform.Get(scene.upgrades.clickerUpgrade2)
	if !found {
		return
	}

	i.Hitbox = &image.Rectangle{
		Min: image.Point{
			X: int(trs.Translation.X) - size.Dx()/2,
			Y: rHeight - (int(trs.Translation.Y) + size.Dy()/2),
		},
		Max: image.Point{
			X: int(trs.Translation.X) + size.Dx()/2,
			Y: rHeight - (int(trs.Translation.Y) - size.Dy()/2),
		},
	}

	i = scene.Components.Interactive.Add(scene.upgrades.clickerUpgrade3)
	i.Callback = func() (preventPropagation bool) {
		scene.upgradeClicker(4)
		scene.RemoveEntity(scene.upgrades.clickerUpgrade3)
		scene.RemoveEntity(scene.upgrades.clickerUpgrade3Label)
		return false
	}
	i.Vector = input.NewInputVector()
	i.Vector.SetMouseButton(input.MouseButtonLeft)
	size, found = scene.Components.Size.Get(scene.upgrades.clickerUpgrade3)
	if !found {
		return
	}

	trs, found = scene.Components.Transform.Get(scene.upgrades.clickerUpgrade3)
	if !found {
		return
	}

	i.Hitbox = &image.Rectangle{
		Min: image.Point{
			X: int(trs.Translation.X) - size.Dx()/2,
			Y: rHeight - (int(trs.Translation.Y) + size.Dy()/2),
		},
		Max: image.Point{
			X: int(trs.Translation.X) + size.Dx()/2,
			Y: rHeight - (int(trs.Translation.Y) - size.Dy()/2),
		},
	}

	i = scene.Components.Interactive.Add(scene.upgrades.clickerUpgrade4)
	i.Callback = func() (preventPropagation bool) {
		scene.upgradeClicker(8)
		scene.RemoveEntity(scene.upgrades.clickerUpgrade4)
		scene.RemoveEntity(scene.upgrades.clickerUpgrade4Label)
		return false
	}
	i.Vector = input.NewInputVector()
	i.Vector.SetMouseButton(input.MouseButtonLeft)
	size, found = scene.Components.Size.Get(scene.upgrades.clickerUpgrade4)
	if !found {
		return
	}

	trs, found = scene.Components.Transform.Get(scene.upgrades.clickerUpgrade4)
	if !found {
		return
	}

	i.Hitbox = &image.Rectangle{
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

func (scene *GameScene) updateBalance(amount int) {
	scene.balanceValue += amount

	balValue, _ := scene.Components.Text.Get(scene.balanceLabel)
	balValue.String = fmt.Sprintf("Balance: " + strconv.Itoa(scene.balanceValue) + " Cubes")
}

func (scene *GameScene) Init(world *akara.World) {
	scene.clickValue = 1
	scene.isDebugEnabled = false
	scene.makeInitialUI()
	scene.bindDebugInput()
	scene.bindClickingInput()
	scene.bindShopClickingInput()
}

func (scene *GameScene) IsInitialized() bool {
	return scene.toggleButton != 0
}
