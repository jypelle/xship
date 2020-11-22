package internal

import (
	"./device"
	"./img"
	"./imgdata"
	"tinygo.org/x/drivers/buzzer"
)

var looseStageTick int
var skullStatus SkullStatus
var looseSong = []float64{
	buzzer.C5,
	buzzer.B4,
	buzzer.A4,
	buzzer.G4,
	buzzer.F4,
}

type SkullStatus struct {
	TickSinceLastRebound int
	V0                   float32
}

func LooseStageReset() {
	looseStageTick = 0
	skullStatus = SkullStatus{TickSinceLastRebound: 27, V0: -210}
}

func LooseStageUpdate() {
	if looseStageTick%6 == 0 && looseStageTick/6 < len(looseSong) {
		device.SoundDevice.Tone(looseSong[looseStageTick/6], 0.012)
	}

	// Update background stars
	starsEntity.Move()

	// Buttons actions
	if device.ButtonsDevice.IsButtonStartPressed() {
		nextGameStage = IntroStage
	}

	looseStageTick++
}

func LooseStageDraw() {
	// Draw stars
	starsEntity.Draw()

	// Draw game over skull
	skullStatus.TickSinceLastRebound++
	t := float32(skullStatus.TickSinceLastRebound) * .032
	y := 56 + int32(skullStatus.V0*t+120.0*t*t)
	if y > 56 {
		skullStatus.TickSinceLastRebound = 0
		skullStatus.V0 = -160
	}
	device.TftDevice.DrawImageFromAsset(imgdata.SkullSprite.ImageAsset, img.Position{X: 64, Y: y}, 0)

	// Draw press start
	if looseStageTick%16 < 8 {
		device.TftDevice.DrawImageFromAsset(imgdata.PressStartSprite.ImageAsset, img.Position{X: 58, Y: 100}, 0)
	}
}
