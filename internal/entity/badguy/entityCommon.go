package badguy

import (
	"github.com/jypelle/xship/internal/img"
)

func NewEntityCommon(position img.Position) entityCommon {
	e := entityCommon{
		state:     BadGuyEnabledState,
		position:  position,
		tickCount: 0,
	}
	return e
}

type entityCommon struct {
	state     State
	position  img.Position
	tickCount int
}

func (e *entityCommon) State() State {
	return e.state
}

func (e *entityCommon) Position() img.Position {
	return e.position
}

func (e *entityCommon) explode() {
	e.tickCount = 0
	e.state = BadGuyExplodingState
}

func (e *entityCommon) Move() {
	e.tickCount++
}
