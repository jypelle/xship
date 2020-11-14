package device

import (
	"machine"
	"tinygo.org/x/drivers/shifter"
)

type buttonsDev struct {
	buttonsShifter shifter.Device
	previousButtonsId, buttonsId uint8
}

func NewButtonsDevice() *buttonsDev{
	return &buttonsDev{
		buttonsShifter: shifter.New(shifter.EIGHT_BITS, machine.BUTTON_LATCH, machine.BUTTON_CLK, machine.BUTTON_OUT),
	}
}

func (d *buttonsDev) configure() {
	d.buttonsShifter.Configure()
}

func (d *buttonsDev) RefreshState() {
	d.previousButtonsId =  d.buttonsId
	d.buttonsId, _ = d.buttonsShifter.Read8Input()
}

func (d *buttonsDev) IsAtLeastOneButtonPressed() bool {
	return d.previousButtonsId == 0 && d.buttonsId != 0
}

func (d *buttonsDev) IsButtonSelectPressed() bool {
	return d.previousButtonsId&machine.BUTTON_SELECT_MASK == 0 && d.buttonsId&machine.BUTTON_SELECT_MASK != 0
}

func (d *buttonsDev) IsButtonStartPressed() bool {
	return d.previousButtonsId&machine.BUTTON_START_MASK == 0 && d.buttonsId&machine.BUTTON_START_MASK != 0
}

func (d *buttonsDev) IsButtonAPressed() bool {
	return d.previousButtonsId&machine.BUTTON_A_MASK == 0 && d.buttonsId&machine.BUTTON_A_MASK != 0
}

func (d *buttonsDev) IsButtonBPressed() bool {
	return d.previousButtonsId&machine.BUTTON_B_MASK == 0 && d.buttonsId&machine.BUTTON_B_MASK != 0
}

func (d *buttonsDev) IsButtonSelectReleased() bool {
	return d.previousButtonsId&machine.BUTTON_SELECT_MASK != 0 && d.buttonsId&machine.BUTTON_SELECT_MASK == 0
}

func (d *buttonsDev) IsButtonStartReleased() bool {
	return d.previousButtonsId&machine.BUTTON_START_MASK != 0 && d.buttonsId&machine.BUTTON_START_MASK == 0
}

func (d *buttonsDev) IsButtonAReleased() bool {
	return d.previousButtonsId&machine.BUTTON_A_MASK != 0 && d.buttonsId&machine.BUTTON_A_MASK == 0
}

func (d *buttonsDev) IsButtonBReleased() bool {
	return d.previousButtonsId&machine.BUTTON_B_MASK != 0 && d.buttonsId&machine.BUTTON_B_MASK == 0
}

func (d *buttonsDev) IsButtonSelectHold() bool {
	return d.buttonsId&machine.BUTTON_SELECT_MASK != 0
}

func (d *buttonsDev) IsButtonStartHold() bool {
	return d.buttonsId&machine.BUTTON_START_MASK != 0
}

func (d *buttonsDev) IsButtonAHold() bool {
	return d.buttonsId&machine.BUTTON_A_MASK != 0
}

func (d *buttonsDev) IsButtonBHold() bool {
	return d.buttonsId&machine.BUTTON_B_MASK != 0
}
