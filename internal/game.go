package internal

import (
	"image/color"

	"github.com/jypelle/xship/internal/device"
	"github.com/jypelle/xship/internal/entity/stars"
	"github.com/jypelle/xship/internal/entity/xship"
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

var xshipEntity *xship.Entity
var starsEntity *stars.Entity
var score int

func Setup() {

	// Setup devices
	device.Setup()

	starsEntity = stars.NewStarsEntity()
}

func Update() {
	device.ButtonsDevice.RefreshState()

	if nextGameStage != gameStage {
		gameStage = nextGameStage

		switch gameStage {
		case IntroStage:
			IntroStageReset()
		case PlayStage:
			PlayStageReset()
		case WinStage:
			WinStageReset()
		case LooseStage:
			LooseStageReset()
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
