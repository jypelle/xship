package internal

import "./img"

type MissileStatus struct {
	Enabled bool
	img.Position
	TickCount int
}
