package stars

import (
	"image/color"
	"math/rand"

	"github.com/jypelle/xship/internal/constant"
	"github.com/jypelle/xship/internal/device"
	"github.com/jypelle/xship/internal/img"
	"github.com/jypelle/xship/internal/imgdata"
	"tinygo.org/x/drivers/st7735"
)

type Entity struct {
	enabled            bool
	intensity          float64
	offsetX            int32
	glowingTick        int32
	glowingId          int32
	blueStarPositions  []img.Position
	whiteStarPositions []img.Position
}

func NewStarsEntity() *Entity {
	s := Entity{
		blueStarPositions:  make([]img.Position, constant.BLUE_STARS_MAX),
		whiteStarPositions: make([]img.Position, constant.WHITE_STARS_MAX),
	}

	// Create starBackground
	for i := 0; i < 50; i++ {
		s.blueStarPositions[i] = img.Position{
			X: rand.Int31n(constant.BOARD_WIDTH),
			Y: rand.Int31n(constant.BOARD_HEIGHT),
		}
	}
	for i := 0; i < 75; i++ {
		s.whiteStarPositions[i] = img.Position{
			X: rand.Int31n(constant.BOARD_WIDTH),
			Y: rand.Int31n(constant.BOARD_HEIGHT),
		}
	}
	s.Reset()
	return &s
}

func (s *Entity) Reset() {
	s.intensity = 0
	s.offsetX = 0
	s.enabled = true
	s.glowingTick = -1
	s.glowingId = -1
}

func (s *Entity) Move() {
	// Move background stars
	s.offsetX = (s.offsetX + 1) % (4 * constant.BOARD_WIDTH)
	s.glowingTick = (s.glowingTick + 1) % (4 * 9)
	if s.glowingTick == 0 {
		s.glowingId = rand.Int31n(constant.WHITE_STARS_MAX)
	}
}

func (s *Entity) Draw() {
	// Draw blue stars
	for i := 0; i < constant.BLUE_STARS_MAX; i++ {
		device.TftDevice.SetPixelFromC565(
			constant.BOARD_XOFFSET+(constant.BOARD_WIDTH+int32(s.blueStarPositions[i].X)-s.offsetX/4)%constant.BOARD_WIDTH,
			constant.BOARD_YOFFSET+int32(s.blueStarPositions[i].Y),
			st7735.RGBATo565(color.RGBA{R: uint8(s.intensity * float64(130)), G: uint8(s.intensity * float64(130)), B: uint8(s.intensity * float64(200)), A: 255}),
		)
	}
	// Draw white stars
	for i := int32(0); i < constant.WHITE_STARS_MAX; i++ {
		if !(i == s.glowingId && s.glowingTick/4 > 0) {
			device.TftDevice.SetPixelFromC565(
				constant.BOARD_XOFFSET+(constant.BOARD_WIDTH*4+s.whiteStarPositions[i].X-s.offsetX/2)%constant.BOARD_WIDTH,
				constant.BOARD_YOFFSET+int32(s.whiteStarPositions[i].Y),
				st7735.RGBATo565(color.RGBA{R: uint8(s.intensity * float64(200)), G: uint8(s.intensity * float64(200)), B: uint8(s.intensity * float64(200)), A: 255}),
			)
		} else {
			device.TftDevice.DrawImageFromAsset(
				imgdata.GlowingWhiteStarSprite.ImageAsset,
				img.Position{
					X: constant.BOARD_XOFFSET + (constant.BOARD_WIDTH*4+s.whiteStarPositions[i].X-s.offsetX/2)%constant.BOARD_WIDTH - 4,
					Y: constant.BOARD_YOFFSET + s.whiteStarPositions[i].Y - 4,
				},
				int(s.glowingTick/4-1),
			)
		}
	}
}

func (s *Entity) Intensity() float64 {
	return s.intensity
}

func (s *Entity) AddIntensity(i float64) {
	s.intensity += i
}
