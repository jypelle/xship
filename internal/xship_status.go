package internal

import "./img"

type XshipStatus struct {
	Position           img.Position
	RemainingBlinkTick int
	HealthPoint        int8
	IsMoving           bool
}

