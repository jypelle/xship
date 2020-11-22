package badguy

import (
	"../../device"
	"../../img"
	"../../imgdata"
	"../missile"
	"../xship"
	"tinygo.org/x/drivers/buzzer"
)

type entity2 struct {
	entityCommon
}

func NewEntity2(position img.Position) *entity2 {
	e := &entity2{
		NewEntityCommon(position),
	}
	return e
}

func (e *entity2) Type() Type {
	return BadGuy2Type
}

func (e *entity2) Move() {
	e.entityCommon.Move()

	if e.state == BadGuyEnabledState {
		e.position.X -= 5
	} else if e.state == BadGuyExplodingState {
		e.position.X -= 2
	}
}

func (e *entity2) Update(xshipEntity *xship.Entity, missileEntities []*missile.Entity, score *int) (newBadGuys []Entity) {
	if e.state == BadGuyEnabledState {
		if e.position.X < -40 {
			e.state = BadGuyDisabledState
			if *score > 0 {
				*score--
			}
		} else {
			// Xship collision
			if xshipEntity.IsVulnerable() && img.IsCollided(&imgdata.XshipSprite.HitBox, xshipEntity.Position(), &imgdata.BadGuy2Sprite.HitBox, e.position) {
				e.explode()
				xshipEntity.Hit(2)
			} else {
				// Missile collision
				for id := range missileEntities {
					if missileEntities[id] != nil {
						if img.IsCollided(&imgdata.MissileSprite.HitBox, missileEntities[id].Position, &imgdata.BadGuy2Sprite.HitBox, e.Position()) {
							e.explode()
							missileEntities[id] = nil
							device.SoundDevice.Tone(buzzer.B1, 0.008)
							*score += 50
						}
					}
				}
			}
		}
	} else if e.state == BadGuyExplodingState {
		if e.tickCount/2 > 2 {
			e.state = BadGuyDisabledState
		}
	}

	return
}

func (e *entity2) Draw() {
	if e.state == BadGuyEnabledState {
		device.TftDevice.DrawImageFromAsset(imgdata.BadGuy2Sprite.ImageAsset, e.position, (e.tickCount/4)%3)
	} else if e.state == BadGuyExplodingState {
		device.TftDevice.DrawImageFromAsset(imgdata.ExplodedBadGuy2Sprite.ImageAsset, e.position, (e.tickCount/2)%3)
	}
}
