package badguy

import (
	"../../device"
	"../../img"
	"../../imgdata"
	"../missile"
	"../xship"
)

type entityMissile struct {
	entityCommon
}

func NewEntityMissile(position img.Position) *entityMissile {
	e := &entityMissile{
		NewEntityCommon(position),
	}
	return e
}

func (e *entityMissile) Type() Type {
	return MissileType
}

func (e *entityMissile) Move() {
	e.entityCommon.Move()
	e.position.X -= 5
}

func (e *entityMissile) Update(xshipEntity *xship.Entity, missileEntities []*missile.Entity, score *int) (newBadGuys []Entity) {
	if e.state == BadGuyEnabledState {
		if e.position.X < -40 {
			e.state = BadGuyDisabledState
		} else {
			// Xship collision
			if xshipEntity.IsVulnerable() && img.IsCollided(&imgdata.XshipSprite.HitBox, xshipEntity.Position(), &imgdata.MissileSprite.HitBox, e.position) {
				e.explode()
				xshipEntity.Hit(1)
			}
		}
	}

	return
}

func (e *entityMissile) Draw() {
	frame := e.tickCount / 2
	if frame >= 2 {
		frame = 2
	}
	device.TftDevice.DrawImageFromAsset(imgdata.BadMissileSprite.ImageAsset, e.position, frame)
}
