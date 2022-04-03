package device

import (
	"machine"
)

type joystickDev struct {
	x machine.ADC
	y machine.ADC
}

func NewJoystickDevice() *joystickDev {
	return &joystickDev{x: machine.ADC{machine.JOYX}, y: machine.ADC{machine.JOYY}}
}

func (d *joystickDev) configure() {
	d.x.Configure(machine.ADCConfig{})
	d.y.Configure(machine.ADCConfig{})
}

func (d *joystickDev) Xaxis() int {
	x := d.x.Get()
	switch {
	case x > 58768:
		return 2
	case x > 38768:
		return 1
	case x > 26768:
		return 0
	case x > 6768:
		return -1
	default:
		return -2
	}
}

func (d *joystickDev) Yaxis() int {
	x := d.y.Get()
	switch {
	case x > 58768:
		return 2
	case x > 38768:
		return 1
	case x > 26768:
		return 0
	case x > 6768:
		return -1
	default:
		return -2
	}
}
