package img

type HitBox struct {
	Offset Position
	Width  int32
	Height int32
}

type Sprite struct {
	ImageAsset
	HitBox HitBox
}

func IsCollided(hitBox1 *HitBox, position1 Position, hitBox2 *HitBox, position2 Position) bool {
	return (position1.X+hitBox1.Offset.X)+hitBox1.Width > position2.X+hitBox2.Offset.X && position1.X+hitBox1.Offset.X < (position2.X+hitBox2.Offset.X+hitBox2.Width) &&
		(position1.Y+hitBox1.Offset.Y)+hitBox1.Height > position2.Y+hitBox2.Offset.Y && position1.Y+hitBox1.Offset.Y < (position2.Y+hitBox2.Height)
}
