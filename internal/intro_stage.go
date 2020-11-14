package internal

import (
	"./constant"
	"./device"
	"./img"
	"image/color"
	"math"
	"tinygo.org/x/drivers/st7735"
)

var introStarIntensity float64

func IntroStageInit() {
	introStarIntensity = 0
	starBackgroundXpos = 0
	xshipStatus.Position = img.Position{X: 0, Y: 56}
	xshipStatus.RemainingBlinkTick = 60
	xshipStatus.HealthPoint = 0
}

func IntroStageUpdate() {
	// Turn on background stars
	introStarIntensity += 1.0 / 60.0

	// Move background stars
	starBackgroundXpos = (starBackgroundXpos + 1) % (4 * constant.BOARD_WIDTH)

	xshipStatus.Position.X++
	if xshipStatus.Position.X >= 60 {
		nextGameStage = PlayStage
	}

	// Increase health point
	xshipStatus.HealthPoint = int8(math.Round(5.0 * introStarIntensity))

	// Update ship
	if xshipStatus.RemainingBlinkTick > 0 {
		xshipStatus.RemainingBlinkTick--
	}
}

func IntroStageDraw() {
	// Refresh Leds
	for i := range neoLeds {
		if int8(i) < xshipStatus.HealthPoint {
			neoLeds[i] = hpColors[i]
		} else {
			neoLeds[i] = blackColor
		}
	}
	device.NeoDevice.WriteColors(neoLeds[:])

	// Draw blue stars
	for i := 0; i < 50; i++ {
		device.TftDevice.SetPixelFromC565(
			constant.BOARD_XOFFSET+(constant.BOARD_WIDTH+int32(blueStarPositions[i].X)-starBackgroundXpos/4)%constant.BOARD_WIDTH,
			constant.BOARD_YOFFSET+int32(blueStarPositions[i].Y),
			st7735.RGBATo565(color.RGBA{R: uint8(introStarIntensity * float64(180)), G: uint8(introStarIntensity * float64(180)), B: uint8(introStarIntensity * float64(255)), A: 255}),
		)
	}
	// Draw white stars
	for i := 0; i < 75; i++ {
		device.TftDevice.SetPixelFromC565(
			constant.BOARD_XOFFSET+(constant.BOARD_WIDTH*4+int32(whiteStarPositions[i].X)-starBackgroundXpos/2)%constant.BOARD_WIDTH,
			constant.BOARD_YOFFSET+int32(whiteStarPositions[i].Y),
			st7735.RGBATo565(color.RGBA{R: uint8(introStarIntensity * float64(255)), G: uint8(introStarIntensity * float64(255)), B: uint8(introStarIntensity * float64(255)), A: 255}),
		)
	}

	// Draw x-ship
	if xshipStatus.RemainingBlinkTick == 0 || xshipStatus.RemainingBlinkTick%4 < 2 {
		device.TftDevice.DrawImageFromAsset(Asset1, XshipSprite.ImageAsset, xshipStatus.Position, 1)
	}
}
