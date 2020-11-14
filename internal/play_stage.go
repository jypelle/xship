package internal

import (
	"./constant"
	"./device"
	"./img"
	"math/rand"
	"tinygo.org/x/drivers/buzzer"
)

var missilesStatus = make([]MissileStatus, constant.MISSILE_MAX)
var missilesIdx = 0
var badGuysStatus = make([]BadGuyStatus, constant.BAD_GUY_MAX)
var badGuysIdx = 0

func PlayStageInit() {
	for id, _ := range badGuysStatus {
		badGuysStatus[id].State = BadGuyDisabledState
		badGuysStatus[id].TickCount = 0
	}

	for id, _ := range missilesStatus {
		missilesStatus[id].Enabled = false
		missilesStatus[id].TickCount = 0
	}
}

func PlayStageUpdate() {

	// Collision actions
	for id, _ := range badGuysStatus {
		badGuysStatus[id].TickCount++

		if badGuysStatus[id].State == BadGuyEnabledState {
			if int32(badGuysStatus[id].X) < -40 {
				badGuysStatus[id].State = BadGuyDisabledState
				badGuysStatus[id].TickCount = 0
			} else {
				if xshipStatus.RemainingBlinkTick == 0 && img.IsCollided(&XshipSprite, xshipStatus.Position, &BadGuySprite, badGuysStatus[id].Position) {
					badGuysStatus[id].State = BadGuyExplodingState
					badGuysStatus[id].TickCount = 0
					xshipStatus.RemainingBlinkTick = 60
					device.SoundDevice.Tone(buzzer.B1, 0.008)
					xshipStatus.HealthPoint--
					if xshipStatus.HealthPoint <= 0 {
						nextGameStage = LooseStage
					}
				}
			}
		} else if badGuysStatus[id].State == BadGuyExplodingState {
			if badGuysStatus[id].TickCount/2 > 2 {
				badGuysStatus[id].State = BadGuyDisabledState
				badGuysStatus[id].TickCount = 0
			}
		}
	}

	for id, _ := range missilesStatus {
		missilesStatus[id].TickCount++

		if missilesStatus[id].Enabled {
			if int32(missilesStatus[id].X) > constant.BOARD_WIDTH {
				missilesStatus[id].Enabled = false
				missilesStatus[id].TickCount = 0
			} else {
				for id2, _ := range badGuysStatus {
					if badGuysStatus[id2].State == BadGuyEnabledState {
						if img.IsCollided(&MissileSprite, missilesStatus[id].Position, &BadGuySprite, badGuysStatus[id2].Position) {
							badGuysStatus[id2].State = BadGuyExplodingState
							badGuysStatus[id2].TickCount = 0
							missilesStatus[id].Enabled = false
							missilesStatus[id].TickCount = 0
							device.SoundDevice.Tone(buzzer.B1, 0.008)
						}
					}
				}
			}
		}
	}

	// Buttons actions
	if device.ButtonsDevice.IsButtonStartPressed() {
		nextGameStage = IntroStage
	}
	if device.ButtonsDevice.IsButtonAPressed() {
		device.SoundDevice.Tone(buzzer.B3, 0.008)
		missilesStatus[missilesIdx] = MissileStatus{
			Position: img.Position{
				X: xshipStatus.Position.X + 22,
				Y: xshipStatus.Position.Y + 9,
			},
			Enabled: true,
		}
		missilesIdx = (missilesIdx + 1) % constant.MISSILE_MAX
	}

	xshipStatus.IsMoving = false
	if device.JoystickDevice.Xaxis() < 0 {
		xshipStatus.IsMoving = true
		xshipStatus.Position.X--
		if device.JoystickDevice.Xaxis() == -2 {
			xshipStatus.Position.X -= 2
		}
		if xshipStatus.Position.X < 1 {
			xshipStatus.Position.X = 1
		}
	}
	if device.JoystickDevice.Xaxis() > 0 {
		xshipStatus.IsMoving = true
		xshipStatus.Position.X++
		if device.JoystickDevice.Xaxis() == 2 {
			xshipStatus.Position.X += 2
		}
		if int32(xshipStatus.Position.X) > constant.BOARD_WIDTH-30 {
			xshipStatus.Position.X = int16(constant.BOARD_WIDTH) - 30
		}
	}
	if device.JoystickDevice.Yaxis() < 0 {
		xshipStatus.IsMoving = true
		xshipStatus.Position.Y--
		if device.JoystickDevice.Yaxis() == -2 {
			xshipStatus.Position.Y -= 2
		}
		if xshipStatus.Position.Y < 0 {
			xshipStatus.Position.Y = 0
		}
	}
	if device.JoystickDevice.Yaxis() > 0 {
		xshipStatus.IsMoving = true
		xshipStatus.Position.Y++
		if device.JoystickDevice.Yaxis() == 2 {
			xshipStatus.Position.Y += 2
		}
		if int32(xshipStatus.Position.Y) > constant.BOARD_HEIGHT-16 {
			xshipStatus.Position.Y = int16(constant.BOARD_HEIGHT) - 16
		}
	}

	// Update ship
	if xshipStatus.RemainingBlinkTick > 0 {
		xshipStatus.RemainingBlinkTick--
	}

	// Move bad guys
	for id, _ := range badGuysStatus {
		if badGuysStatus[id].State == BadGuyEnabledState {
			badGuysStatus[id].X -= 3
		} else if badGuysStatus[id].State == BadGuyExplodingState {
			badGuysStatus[id].X -= 2
		}
	}

	// Move active missilesStatus
	for id, _ := range missilesStatus {
		if missilesStatus[id].Enabled {
			missilesStatus[id].X += 3
		}
	}

	// Move background stars
	starBackgroundXpos = (starBackgroundXpos + 1) % (4 * constant.BOARD_WIDTH)

	// Generate new bad guy
	if rand.Intn(8) == 0 {
		badGuysStatus[badGuysIdx] = BadGuyStatus{
			Position: img.Position{
				X: int16(constant.BOARD_WIDTH),
				Y: int16(rand.Intn(int(constant.BOARD_HEIGHT - BadGuySprite.Height))),
			},
			State: BadGuyEnabledState,
		}
		badGuysIdx = (badGuysIdx + 1) % constant.BAD_GUY_MAX
	}
}

func PlayStageDraw() {

	// Refresh health bar
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
			blueStarC565,
		)
	}

	// Draw white stars
	for i := 0; i < 75; i++ {
		device.TftDevice.SetPixelFromC565(
			constant.BOARD_XOFFSET+(constant.BOARD_WIDTH*4+int32(whiteStarPositions[i].X)-starBackgroundXpos/2)%constant.BOARD_WIDTH,
			constant.BOARD_YOFFSET+int32(whiteStarPositions[i].Y),
			whiteStarC565,
		)
	}

	// Draw missiles
	for _, missile := range missilesStatus {
		if missile.Enabled {
			device.TftDevice.DrawImageFromAsset(Asset1, MissileSprite.ImageAsset, missile.Position, 0)
		}
	}

	// Draw x-ship
	if xshipStatus.RemainingBlinkTick == 0 || xshipStatus.RemainingBlinkTick%4 < 2 {
		if xshipStatus.IsMoving {
			device.TftDevice.DrawImageFromAsset(Asset1, XshipSprite.ImageAsset, xshipStatus.Position, 1)
		} else {
			device.TftDevice.DrawImageFromAsset(Asset1, XshipSprite.ImageAsset, xshipStatus.Position, 0)
		}
	}

	// Draw bad guys
	for _, badGuy := range badGuysStatus {
		if badGuy.State == BadGuyEnabledState {
			device.TftDevice.DrawImageFromAsset(Asset1, BadGuySprite.ImageAsset, badGuy.Position, (badGuy.TickCount/4)%3)
		} else if badGuy.State == BadGuyExplodingState {
			device.TftDevice.DrawImageFromAsset(Asset1, ExplodedBadGuySprite.ImageAsset, badGuy.Position, (badGuy.TickCount/2)%3)
		}
	}

}
