package units

import "rpg/game/maps"

type Unit struct {
	ID                  string           `json:"id"`
	MyCoordinate        *maps.Coordinate `json:"myCooord"`
	Road                *maps.Coordinate `json:"target"`
	SpriteName          string           `json:"sprite_name"`
	Action              string           `json:"action"`
	Frame               int              `json:"frame"`
	HorizontalDirection DirectionType    `json:"horizontal"`
}

// type Units map[string]IsUnit

type IsUnit interface {
	GET() *Unit
	Create() IsUnit
	CreateInUnits(un Unit) IsUnit
	GetID() string
	UpdateCoordinate(int, int)
	GetCoordinate() (float64, float64)
	UpdateAction(action string)
	GetAction() string
	UpdateRoad(road *maps.Coordinate)
	GetRoad() *maps.Coordinate
	Unmarshal([]byte) (map[string]IsUnit, error)
}

type DirectionType int

const (
	DirectionUp    DirectionType = 0
	DirectionDown  DirectionType = 1
	DirectionLeft  DirectionType = 2
	DirectionRight DirectionType = 3
)

const (
	ActionRun  = "run"
	ActionIdle = "idle"
)
