package badguy

import (
	"../../device"
	"../../img"
	"../../imgdata"
	"../missile"
	"../xship"
	"tinygo.org/x/drivers/buzzer"
)

type entity1 struct {
	entityCommon
	yref int32
}

func NewEntity1(position img.Position) *entity1 {
	e := &entity1{
		entityCommon: NewEntityCommon(position),
		yref:         position.Y,
	}
	return e
}

func (e *entity1) Type() Type {
	return BadGuy1Type
}

func (e *entity1) Move() {
	e.entityCommon.Move()

	if e.state == BadGuyEnabledState {
		e.position.X -= 2
		//		e.position.Y = e.yref + int32(5.0*math.Cos(float64(e.tickCount)/3.0))
	} else if e.state == BadGuyExplodingState {
		e.position.X -= 2
	}
}

func (e *entity1) Update(xshipEntity *xship.Entity, missileEntities []*missile.Entity, score *int) (newBadGuys []Entity) {
	if e.state == BadGuyEnabledState {
		if e.position.X < -40 {
			e.state = BadGuyDisabledState
			if *score > 0 {
				*score--
			}
		} else {
			// Xship collision
			if xshipEntity.IsVulnerable() && img.IsCollided(&imgdata.XshipSprite.HitBox, xshipEntity.Position(), &imgdata.BadGuy1Sprite.HitBox, e.position) {
				e.explode()
				xshipEntity.Hit(1)
			} else {
				// Missile collision
				for id := range missileEntities {
					if missileEntities[id] != nil {
						if img.IsCollided(&imgdata.MissileSprite.HitBox, missileEntities[id].Position, &imgdata.BadGuy1Sprite.HitBox, e.Position()) {
							e.explode()
							missileEntities[id] = nil
							device.SoundDevice.Tone(buzzer.B1, 0.008)
							*score += 5
						}
					}
				}

				// BadGuy shoot a BadGuyMissile
				if (e.tickCount-1)%4 == 0 && ((e.tickCount-1)/16)%3 == 0 {
					newBadGuys = append(newBadGuys, NewEntityMissile(img.Position{
						X: e.position.X - 2,
						Y: e.position.Y + 7,
					}))
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

func (e *entity1) Draw() {
	if e.state == BadGuyEnabledState {
		device.TftDevice.DrawImageFromAsset(imgdata.BadGuy1Sprite.ImageAsset, e.position, (e.tickCount/4)%3)
	} else if e.state == BadGuyExplodingState {
		device.TftDevice.DrawImageFromAsset(imgdata.ExplodedBadGuy1Sprite.ImageAsset, e.position, (e.tickCount/2)%3)
	}
}
