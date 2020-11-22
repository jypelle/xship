package internal

import (
	"./device"
	"./entity/xship"
	"math"
)

func IntroStageReset() {
	starsEntity.Reset()
	xshipEntity = xship.NewEntity()
}

func IntroStageUpdate() {
	// Increase background stars intensity
	starsEntity.AddIntensity(1.0 / 60.0)

	// Move background stars
	starsEntity.Move()

	// Move ship
	xshipEntity.IntroMove()
	if xshipEntity.Position().X >= 60 {
		nextGameStage = PlayStage
	}

	// Increase health point
	xshipEntity.SetHealthPoint(int(math.Round(5.0 * starsEntity.Intensity())))

	// Update ship
	xshipEntity.IntroUpdate()
}

func IntroStageDraw() {
	// Refresh Leds
	for i := range neoLeds {
		if i < xshipEntity.HealthPoint() {
			neoLeds[i] = hpColors[i]
		} else {
			neoLeds[i] = blackColor
		}
	}
	device.NeoDevice.WriteColors(neoLeds[:])

	// Draw stars
	starsEntity.Draw()

	// Draw x-ship
	xshipEntity.Draw()
}
