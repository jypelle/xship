package xship

import (
	"github.com/jypelle/xship/internal/constant"
	"github.com/jypelle/xship/internal/device"
	"github.com/jypelle/xship/internal/entity/missile"
	"github.com/jypelle/xship/internal/img"
	"github.com/jypelle/xship/internal/imgdata"
	"tinygo.org/x/drivers/buzzer"
)

type Entity struct {
	position           img.Position
	remainingBlinkTick int
	healthPoint        int
	isMoving           bool
	autoFire           bool
	autoFireTick       int
}

func NewEntity() *Entity {
	e := Entity{
		position:           img.Position{X: 0, Y: 56},
		remainingBlinkTick: 60,
	}
	return &e
}

func (e *Entity) Position() img.Position {
	return e.position
}

func (e *Entity) IsVulnerable() bool {
	return e.remainingBlinkTick == 0
}

func (e *Entity) IntroMove() {
	e.isMoving = true
	e.position.X++
}

func (e *Entity) Move() {
	e.isMoving = false
	if device.JoystickDevice.Xaxis() < 0 {
		e.isMoving = true
		e.position.X--
		if device.JoystickDevice.Xaxis() == -2 {
			e.position.X -= 2
		}
		if e.position.X < 0 {
			e.position.X = 0
		}
	}
	if device.JoystickDevice.Xaxis() > 0 {
		e.isMoving = true
		e.position.X++
		if device.JoystickDevice.Xaxis() == 2 {
			e.position.X += 3
		}
		if e.position.X > constant.BOARD_WIDTH-30 {
			e.position.X = constant.BOARD_WIDTH - 30
		}
	}
	if device.JoystickDevice.Yaxis() < 0 {
		e.isMoving = true
		e.position.Y--
		if device.JoystickDevice.Yaxis() == -2 {
			e.position.Y -= 3
		}
		if e.position.Y < 0 {
			e.position.Y = 0
		}
	}
	if device.JoystickDevice.Yaxis() > 0 {
		e.isMoving = true
		e.position.Y++
		if device.JoystickDevice.Yaxis() == 2 {
			e.position.Y += 3
		}
		if e.position.Y > constant.BOARD_HEIGHT-16 {
			e.position.Y = constant.BOARD_HEIGHT - 16
		}
	}
	if e.autoFire {
		e.autoFireTick++
	}

}

func (e *Entity) IntroUpdate() {
	if e.remainingBlinkTick > 0 {
		e.remainingBlinkTick--
	}
}

func (e *Entity) Update(missilesIdx *int, missileEntities []*missile.Entity) {
	// Enable/disable autofire
	if device.ButtonsDevice.IsButtonAPressed() {
		e.autoFire = true
		e.autoFireTick = 0
	}
	if device.ButtonsDevice.IsButtonAReleased() {
		e.autoFire = false
	}

	// Generate 1 missile each 4 ticks when autofire is enabled
	if e.autoFire && e.autoFireTick%4 == 0 {
		device.SoundDevice.Tone(buzzer.B3, 0.008)
		missileEntities[*missilesIdx] = &missile.Entity{
			Position: img.Position{
				X: e.position.X + 22,
				Y: e.position.Y + 9,
			},
		}
		*missilesIdx = (*missilesIdx + 1) % constant.MISSILE_ENTITY_MAX
	}

	if e.remainingBlinkTick > 0 {
		e.remainingBlinkTick--
	}
}

func (e *Entity) Draw() {
	if e.remainingBlinkTick == 0 || e.remainingBlinkTick%4 < 2 {
		if e.isMoving {
			device.TftDevice.DrawImageFromAsset(imgdata.XshipSprite.ImageAsset, e.position, 1)
		} else {
			device.TftDevice.DrawImageFromAsset(imgdata.XshipSprite.ImageAsset, e.position, 0)
		}
	}
}

func (e *Entity) Hit(hitPoint int) {
	e.remainingBlinkTick = 60
	device.SoundDevice.Tone(buzzer.B1, 0.008)
	e.healthPoint -= hitPoint
}

func (e *Entity) SetHealthPoint(hp int) {
	e.healthPoint = hp
	if e.healthPoint > 5 {
		e.healthPoint = 5
	}
}

func (e *Entity) HealthPoint() int {
	return e.healthPoint
}
