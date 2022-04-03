package internal

import (
	"math/rand"

	"github.com/jypelle/xship/internal/constant"
	"github.com/jypelle/xship/internal/device"
	"github.com/jypelle/xship/internal/entity/badguy"
	"github.com/jypelle/xship/internal/entity/missile"
	"github.com/jypelle/xship/internal/img"
	"github.com/jypelle/xship/internal/imgdata"
)

var missileEntities = make([]*missile.Entity, constant.MISSILE_ENTITY_MAX)
var missilesIdx int
var badGuyEntities = make([]badguy.Entity, 0, constant.BAD_GUY_ENTITY_MAX)
var playStagePause bool

func PlayStageReset() {

	playStagePause = false

	// Reset score
	score = 0

	// Clean bad guys
	badGuyEntities = nil

	// Clean missiles
	missilesIdx = 0
	for id := range missileEntities {
		missileEntities[id] = nil
	}
}

func PlayStageUpdate() {

	if playStagePause {
		if device.ButtonsDevice.IsButtonStartPressed() {
			playStagePause = false
		}
	} else {
		// Move background stars
		starsEntity.Move()

		// Move badguys
		for _, badGuyEntity := range badGuyEntities {
			if badGuyEntity != nil {
				badGuyEntity.Move()
			}
		}

		// Buttons actions
		if device.ButtonsDevice.IsButtonSelectPressed() {
			nextGameStage = IntroStage
		}
		if device.ButtonsDevice.IsButtonStartPressed() {
			playStagePause = true
		}

		// Move ship
		xshipEntity.Move()

		// Move active missileEntities
		for _, missileEntity := range missileEntities {
			if missileEntity != nil {
				missileEntity.Move()
			}
		}

		// Update ship
		xshipEntity.Update(&missilesIdx, missileEntities)
		if xshipEntity.HealthPoint() <= 0 {
			nextGameStage = LooseStage
		}

		// Update missiles
		for id := range missileEntities {
			if missileEntities[id] != nil {
				if missileEntities[id].X > constant.BOARD_WIDTH {
					missileEntities[id] = nil
				}
			}
		}

		// Update badguys
		newId := 0
		var newBadGuys []badguy.Entity
		for id, badGuyEntity := range badGuyEntities {
			newBadGuys = append(newBadGuys, badGuyEntity.Update(xshipEntity, missileEntities, &score)...)
			// Remove disabled badGuy
			if badGuyEntity.State() == badguy.BadGuyDisabledState {
				badGuyEntities[id] = nil
			} else {
				badGuyEntities[newId] = badGuyEntity
				newId++
			}
		}
		badGuyEntities = badGuyEntities[:newId]
		badGuyEntities = append(badGuyEntities, newBadGuys...)

		// Generate new bad guy
		if rand.Intn(30) == 0 {
			badGuyEntities = append(badGuyEntities, badguy.NewEntity1(img.Position{
				X: constant.BOARD_WIDTH,
				Y: rand.Int31n(constant.BOARD_HEIGHT - imgdata.BadGuy1Sprite.Height),
			}))
		}
		if rand.Intn(140) == 0 {
			badGuyEntities = append(badGuyEntities, badguy.NewEntity2(img.Position{
				X: constant.BOARD_WIDTH,
				Y: rand.Int31n(constant.BOARD_HEIGHT - imgdata.BadGuy2Sprite.Height),
			}))
		}

	}
}

func PlayStageDraw() {

	// Refresh health bar
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

	// Draw missiles
	for _, missile := range missileEntities {
		if missile != nil {
			missile.Draw()
		}
	}

	// Draw x-ship
	xshipEntity.Draw()

	// Draw bad guys
	for _, badGuy := range badGuyEntities {
		if badGuy != nil {
			badGuy.Draw()
		}
	}

	// Draw score
	device.TftDevice.DrawImageFromAsset(imgdata.NumberSprite.ImageAsset, img.Position{X: device.TftDevice.Width() - 3}, score%10)
	device.TftDevice.DrawImageFromAsset(imgdata.NumberSprite.ImageAsset, img.Position{X: device.TftDevice.Width() - 7}, score/10%10)
	device.TftDevice.DrawImageFromAsset(imgdata.NumberSprite.ImageAsset, img.Position{X: device.TftDevice.Width() - 11}, score/100%10)
	device.TftDevice.DrawImageFromAsset(imgdata.NumberSprite.ImageAsset, img.Position{X: device.TftDevice.Width() - 15}, score/1000%10)
	device.TftDevice.DrawImageFromAsset(imgdata.NumberSprite.ImageAsset, img.Position{X: device.TftDevice.Width() - 19}, score/10000%10)
	device.TftDevice.DrawImageFromAsset(imgdata.NumberSprite.ImageAsset, img.Position{X: device.TftDevice.Width() - 23}, score/100000%10)

	// Draw pause
	if playStagePause {
		device.TftDevice.DrawImageFromAsset(imgdata.PauseSprite.ImageAsset, img.Position{X: 69, Y: 60}, 0)
	}

}
