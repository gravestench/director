package main

import (
	"fmt"
	"image"
	"image/color"
<<<<<<< HEAD
	"strconv"
=======
>>>>>>> af206846978268e8fbd85243a19b2ef9b5cbb2ee

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/systems/input"

	"github.com/gravestench/director/pkg/systems/scene"
)

type GameScene struct {
	scene.Scene
<<<<<<< HEAD
	toggleButton    akara.EID
	toggleLabel     akara.EID
	debugPanel      akara.EID
	mainPanel       akara.EID
	square          akara.EID
	label           akara.EID
	mouseDebugLabel akara.EID
	isDebugEnabled  bool
	balanceLabel    akara.EID
	balanceValue    int
	clickButton     akara.EID
	clickValue      int
	shopPanel       akara.EID
	upgrades        ShopUpgrades
}

type ShopUpgrades struct {
	clickerUpgrade1      akara.EID
	clickerUpgrade1Label akara.EID
	clickerUpgrade1Price int
	clickerUpgrade2      akara.EID
	clickerUpgrade2Label akara.EID
	clickerUpgrade2Price int
	clickerUpgrade3      akara.EID
	clickerUpgrade3Label akara.EID
	clickerUpgrade3Price int
	clickerUpgrade4      akara.EID
	clickerUpgrade4Label akara.EID
	clickerUpgrade4Price int
=======
	toggleButton   akara.EID
	toggleLabel    akara.EID
	debugPanel     akara.EID
	mainPanel      akara.EID
	square         akara.EID
	label          akara.EID
	testLabel      akara.EID
	isDebugEnabled bool
>>>>>>> af206846978268e8fbd85243a19b2ef9b5cbb2ee
}

// Game Loop
func (scene *GameScene) Update() {
<<<<<<< HEAD

	//scene.updateLabel()
	if scene.isDebugEnabled == false {

	} else {
		scene.updateTestLabel()
	}
=======
	scene.updateTestLabel()
	//scene.updateLabel()

>>>>>>> af206846978268e8fbd85243a19b2ef9b5cbb2ee
}

func (scene *GameScene) makeMainPanel() {
	background := color.RGBA{R: 30, G: 31, B: 35, A: 255}
	scene.toggleButton = scene.Add.Rectangle(scene.Window.Width/2, scene.Window.Height/2, scene.Window.Width, scene.Window.Height, background, nil)
}

/****************************
*	  Toggle debug code		*
****************************/
<<<<<<< HEAD
=======

>>>>>>> af206846978268e8fbd85243a19b2ef9b5cbb2ee
func (scene *GameScene) makeToggleButton() {
	purple := color.RGBA{R: 104, G: 70, B: 236, A: 255}
	scene.toggleButton = scene.Add.Rectangle(scene.Window.Width-60, scene.Window.Height-15, 140, 30, purple, nil)
}
<<<<<<< HEAD
func (scene *GameScene) toggleDebug() {
	if scene.isDebugEnabled == false {
		scene.isDebugEnabled = true
		scene.makeDebugPanel()
		scene.makeMouseDebugLabel()
	} else {
		scene.isDebugEnabled = false
		scene.clearDebugPanel()
	}
}
func (scene *GameScene) clearDebugPanel() {
	scene.Director.RemoveEntity(scene.debugPanel)
	scene.Director.RemoveEntity(scene.mouseDebugLabel)
}
func (scene *GameScene) makeDebugPanel() {
	background := color.RGBA{R: 21, G: 23, B: 24, A: 255}
	scene.debugPanel = scene.Add.Rectangle(scene.Window.Width/2-132, scene.Window.Height-30, scene.Window.Width, 60, background, nil)
}
func (scene *GameScene) makeToggleLabel() {
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	scene.toggleLabel = scene.Add.Label("Toggle Debug", scene.Window.Width-30, scene.Window.Height-15, 12, "", white)
}
func (scene *GameScene) makeMouseDebugLabel() {
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	scene.mouseDebugLabel = scene.Add.Label("Mouse: ", scene.Window.Width-scene.Window.Width, scene.Window.Height-15, 12, "", white)
	origin, found := scene.Components.Origin.Get(scene.mouseDebugLabel)
	if !found {
		return
	}
	origin.X = 0
	origin.Y = 0
}
=======

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

>>>>>>> af206846978268e8fbd85243a19b2ef9b5cbb2ee
func (scene *GameScene) bindDebugInput() {
	i := scene.Components.Interactive.Add(scene.toggleButton)

	i.Callback = func() (preventPropogation bool) {
		scene.toggleDebug()
<<<<<<< HEAD
=======
		fmt.Print("Test")
>>>>>>> af206846978268e8fbd85243a19b2ef9b5cbb2ee
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
<<<<<<< HEAD
func (scene *GameScene) updateTestLabel() {
	text, found := scene.Components.Text.Get(scene.mouseDebugLabel)
=======

/****************************
*	End toggle debug code	*
****************************/

func (scene *GameScene) updateTestLabel() {
	text, found := scene.Components.Text.Get(scene.testLabel)
>>>>>>> af206846978268e8fbd85243a19b2ef9b5cbb2ee
	if !found {
		return
	}

	mp := rl.GetMousePosition()

	const (
<<<<<<< HEAD
		fmtMouse = "Mouse: (%v, %v)"
=======
		fmtMouse = "Mouse (%v, %v)"
>>>>>>> af206846978268e8fbd85243a19b2ef9b5cbb2ee
	)

	text.String = fmt.Sprintf(fmtMouse, mp.X, mp.Y)
}

<<<<<<< HEAD
/****************************
*	End toggle debug code	*
****************************/

func (scene *GameScene) makeInitialUI() {
	gray := color.RGBA{R: 97, G: 98, B: 109, A: 255}

	scene.makeMainPanel()
	scene.makeToggleButton()
	scene.makeToggleLabel()
	scene.makeShopPanel()
	scene.balanceLabel = scene.Add.Label("Balance: 0 Cubes", scene.Window.Width/2+40, scene.Window.Height/2+85, 24, "", gray)
	// origin, found := scene.Components.Origin.Get(scene.balanceLabel)
	// if !found {
	// 	return
	// }
	// origin.X = 0
	// origin.Y = 0
	scene.makeClickButton()
}

func (scene *GameScene) makeClickButton() {
	purple := color.RGBA{R: 104, G: 70, B: 236, A: 255}
	scene.clickButton = scene.Add.Rectangle(scene.Window.Width/2, scene.Window.Height/2, 140, 140, purple, nil)
}

func (scene *GameScene) makeShopPanel() {
	dark := color.RGBA{R: 13, G: 12, B: 15, A: 255}
	shopWidth := 150
	scene.shopPanel = scene.Add.Rectangle(shopWidth/2, scene.Window.Height/2, shopWidth, scene.Window.Height, dark, nil)
	scene.makeShopUpgrades()
}

func (scene *GameScene) makeShopUpgrades() {
	upgradeYLocation := scene.Window.Height - 20
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
	upgradeYLocation -= upgradeSize
}

func (scene *GameScene) upgradeClicker(value int) {
	scene.clickValue += value
}

func (scene *GameScene) bindClickingInput() {
	i := scene.Components.Interactive.Add(scene.clickButton)
	i.Callback = func() (preventPropogation bool) {
		scene.updateBalance(scene.clickValue)
		return false
	}
	i.Vector = input.NewInputVector()
	i.Vector.SetMouseButton(input.MouseButtonLeft)
	size, found := scene.Components.Size.Get(scene.clickButton)
	if !found {
		return
	}

	trs, found := scene.Components.Transform.Get(scene.clickButton)
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
}

func (scene *GameScene) bindShopClickingInput() {
	i := scene.Components.Interactive.Add(scene.upgrades.clickerUpgrade1)
	i.Callback = func() (preventPropogation bool) {
		fmt.Print(scene.clickValue)
		scene.upgradeClicker(1)
		scene.Director.RemoveEntity(scene.upgrades.clickerUpgrade1)
		scene.Director.RemoveEntity(scene.upgrades.clickerUpgrade1Label)
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

	i = scene.Components.Interactive.Add(scene.upgrades.clickerUpgrade2)
	i.Callback = func() (preventPropogation bool) {
		scene.upgradeClicker(2)
		scene.Director.RemoveEntity(scene.upgrades.clickerUpgrade2)
		scene.Director.RemoveEntity(scene.upgrades.clickerUpgrade2Label)
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

	i = scene.Components.Interactive.Add(scene.upgrades.clickerUpgrade3)
	i.Callback = func() (preventPropogation bool) {
		scene.upgradeClicker(4)
		scene.Director.RemoveEntity(scene.upgrades.clickerUpgrade3)
		scene.Director.RemoveEntity(scene.upgrades.clickerUpgrade3Label)
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

	i = scene.Components.Interactive.Add(scene.upgrades.clickerUpgrade4)
	i.Callback = func() (preventPropogation bool) {
		scene.upgradeClicker(8)
		scene.Director.RemoveEntity(scene.upgrades.clickerUpgrade4)
		scene.Director.RemoveEntity(scene.upgrades.clickerUpgrade4Label)
		return false
	}
	i.Vector = input.NewInputVector()
	i.Vector.SetMouseButton(input.MouseButtonLeft)
	size, found = scene.Components.Size.Get(scene.upgrades.clickerUpgrade4)
	if !found {
		return
	}

	trs, found = scene.Components.Transform.Get(scene.upgrades.clickerUpgrade4)
=======
func (scene *GameScene) updateLabel() {
	text, found := scene.Components.Text.Get(scene.label)
>>>>>>> af206846978268e8fbd85243a19b2ef9b5cbb2ee
	if !found {
		return
	}

<<<<<<< HEAD
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

}

func (scene *GameScene) updateBalance(amount int) {
	scene.balanceValue += amount

	balValue, found := scene.Components.Text.Get(scene.balanceLabel)
	if !found {
		return
	}

	balValue.String = fmt.Sprintf("Balance: " + strconv.Itoa(scene.balanceValue) + " Cubes")
}

func (scene *GameScene) Init(world *akara.World) {
	scene.clickValue = 1
	scene.isDebugEnabled = false
	scene.makeInitialUI()
	scene.bindDebugInput()
	scene.bindClickingInput()
	scene.bindShopClickingInput()
=======
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
>>>>>>> af206846978268e8fbd85243a19b2ef9b5cbb2ee
}

func (scene *GameScene) makeSquare() {
	blue := color.RGBA{B: 255, A: 255}
	scene.square = scene.Add.Rectangle(100, 100, 30, 30, blue, nil)
}

<<<<<<< HEAD
func (scene *GameScene) IsInitialized() bool {
=======
func (scene *GameScene) makeLabel() {
	red := color.RGBA{R: 255, A: 255}
	scene.label = scene.Add.Label("", 400, 400, 24, "", red)
	scene.testLabel = scene.Add.Label("", 200, 200, 24, "", red)
}

func (scene *GameScene) IsInitialized() bool {

>>>>>>> af206846978268e8fbd85243a19b2ef9b5cbb2ee
	return scene.toggleButton != 0
}
