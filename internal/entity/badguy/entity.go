package badguy

import (
	"github.com/jypelle/xship/internal/entity/missile"
	"github.com/jypelle/xship/internal/entity/xship"
	"github.com/jypelle/xship/internal/img"
)

type Type uint8

const (
	BadGuy1Type Type = iota
	BadGuy2Type
	MissileType
)

type State int8

const (
	BadGuyDisabledState State = iota
	BadGuyExplodingState
	BadGuyEnabledState
)

type Entity interface {
	Type() Type
	State() State
	Position() img.Position

	Move()
	Update(xshipEntity *xship.Entity, missileEntities []*missile.Entity, score *int) (newBadGuys []Entity)
	Draw()
}
