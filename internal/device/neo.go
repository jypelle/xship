package device

import (
	"machine"

	"tinygo.org/x/drivers/ws2812"
)

type neoDev struct {
	ws2812.Device
}

func NewNeoDevice(pin machine.Pin) *neoDev {
	return &neoDev{
		Device: ws2812.New(pin),
	}
}

func (d *neoDev) configure() {
	d.Pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
}
