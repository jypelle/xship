package missile

import (
	"github.com/jypelle/xship/internal/device"
	"github.com/jypelle/xship/internal/img"
	"github.com/jypelle/xship/internal/imgdata"
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
