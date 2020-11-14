package device

import (
	"machine"
	"tinygo.org/x/drivers/buzzer"
)

type soundDev struct {
	buzzer.Device
}

func NewSoundDevice(pin machine.Pin) *soundDev{
	return &soundDev{
		Device: buzzer.New(pin),
	}
}

func (d *soundDev) configure() {
	speaker := machine.SPEAKER_ENABLE
	speaker.Configure(machine.PinConfig{Mode: machine.PinOutput})
	speaker.Set(true)

	machine.A0.Configure(machine.PinConfig{Mode: machine.PinOutput})
}