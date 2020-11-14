package internal

import (
	"./device"
	"./img"
)

var looseStageTick int

func LooseStageInit() {
	looseStageTick = 0
}

func LooseStageUpdate() {
	looseStageTick++

	// Buttons actions
	if device.ButtonsDevice.IsButtonStartPressed() {
		nextGameStage = IntroStage
	}
}

func LooseStageDraw() {
	// Draw game over skull
	device.TftDevice.DrawImageFromAsset(Asset1, SkullSprite.ImageAsset, img.Position{X: 64, Y: 38}, 0)

	// Draw press start
	if looseStageTick%16 < 8 {
		device.TftDevice.DrawImageFromAsset(Asset1, PressStartSprite.ImageAsset, img.Position{X: 58, Y: 100}, 0)
	}
}
