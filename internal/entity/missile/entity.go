package missile

import (
	"../../device"
	"../../img"
	"../../imgdata"
)

type Entity struct {
	img.Position
	tickCount int
}

func (e *Entity) Move() {
	e.tickCount++
	e.X += 5
}

func (e *Entity) Draw() {
	device.TftDevice.DrawImageFromAsset(imgdata.MissileSprite.ImageAsset, e.Position, 0)
}
