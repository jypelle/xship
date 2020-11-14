package img

type HitBox struct {
	Offset Position
	Width  int32
	Height int32
}

type Sprite struct {
	ImageAsset
	HitBox   HitBox
}

func IsCollided(sprite1 *Sprite, position1 Position, sprite2 *Sprite, position2 Position) bool {
	return (int32(position1.X + sprite1.HitBox.Offset.X) + sprite1.HitBox.Width) > int32(position2.X + sprite2.HitBox.Offset.X) && int32(position1.X + sprite1.HitBox.Offset.X) < (int32(position2.X + sprite2.HitBox.Offset.X) + sprite2.HitBox.Width) &&
	 (int32(position1.Y + sprite1.HitBox.Offset.Y) + sprite1.HitBox.Height) > int32(position2.Y + sprite2.HitBox.Offset.Y) && int32(position1.Y + sprite1.HitBox.Offset.Y) < (int32(position2.Y) + sprite2.HitBox.Height)
}
