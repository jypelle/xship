package internal

import (
	"./constant"
	"./device"
	"./img"
	"image/color"
	"math/rand"
	"tinygo.org/x/drivers/st7735"
)

// Buffers
var neoLeds [5]color.RGBA

var hpColors = []color.RGBA{
	{R: 8, G: 0, B: 0, A: 255},
	{R: 6, G: 2, B: 0, A: 255},
	{R: 4, G: 4, B: 0, A: 255},
	{R: 2, G: 6, B: 0, A: 255},
	{R: 0, G: 8, B: 0, A: 255},
}
var blackColor = color.RGBA{R: 0, G: 0, B: 0, A: 255}
var blueStarPositions = make([]img.Position, 50)
var whiteStarPositions = make([]img.Position, 75)
var whiteStarC565 = st7735.RGBATo565(color.RGBA{R: 255, G: 255, B: 255, A: 255})
var blueStarC565 = st7735.RGBATo565(color.RGBA{R: 180, G: 180, B: 255, A: 255})

type GameStage int

const (
	UndefinedStage GameStage = iota
	IntroStage
	PlayStage
	WinStage
	LooseStage
)

var gameStage = UndefinedStage
var nextGameStage = IntroStage

var starBackgroundXpos int32
var xshipStatus XshipStatus

func Setup() {

	// Setup devices
	device.Setup()

	// Create starBackground
	for i := 0; i < 50; i++ {
		blueStarPositions[i] = img.Position{
			X: int16(rand.Intn(int(constant.BOARD_WIDTH))),
			Y: int16(rand.Intn(int(constant.BOARD_HEIGHT))),
		}
	}
	for i := 0; i < 75; i++ {
		whiteStarPositions[i] = img.Position{
			X: int16(rand.Intn(int(constant.BOARD_WIDTH))),
			Y: int16(rand.Intn(int(constant.BOARD_HEIGHT))),
		}
	}

}

func Update() {
	device.ButtonsDevice.RefreshState()

	if nextGameStage != gameStage {
		gameStage = nextGameStage

		switch gameStage {
		case IntroStage:
			IntroStageInit()
		case PlayStage:
			PlayStageInit()
		case WinStage:
			WinStageInit()
		case LooseStage:
			LooseStageInit()
		}
	}

	switch gameStage {
	case IntroStage:
		IntroStageUpdate()
	case PlayStage:
		PlayStageUpdate()
	case WinStage:
		WinStageUpdate()
	case LooseStage:
		LooseStageUpdate()
	}

}

func Draw() {

	// Clear Screen
	device.TftDevice.Clear()

	switch gameStage {
	case IntroStage:
		IntroStageDraw()
	case PlayStage:
		PlayStageDraw()
	case WinStage:
		WinStageDraw()
	case LooseStage:
		LooseStageDraw()
	}

	// Send image to tftDevice screen
	device.TftDevice.Refresh()

}
