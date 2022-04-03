package device

import (
	"machine"

	"github.com/jypelle/xship/internal/constant"
)

var NeoDevice = NewNeoDevice(machine.NEOPIXELS)
var ButtonsDevice = NewButtonsDevice()
var JoystickDevice = NewJoystickDevice()
var TftDevice = NewTftDevice(constant.SCREEN_WIDTH, constant.SCREEN_HEIGHT)
var SoundDevice = NewSoundDevice(machine.A0)

func Setup() {
	machine.InitADC()

	// Neopixel
	NeoDevice.configure()

	// Buttons
	ButtonsDevice.configure()

	// Joystick
	JoystickDevice.configure()

	// Tft
	TftDevice.configure()

	// Sound
	SoundDevice.configure()
}
