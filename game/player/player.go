package player

import (
	"encoding/json"
	"log"
	"math/rand"

	"rpg/game/maps"
	"rpg/game/units"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Player struct {
	units.Unit `json:"unit"`
	Weapon     string `json:"weapon"`
	Armor      string `json:"armor"`
	Boots      string `json:"boots"`
}

type Players map[string]*Player

func (p *Player) Create() units.IsUnit {
	skins := []string{
		"elf_f", "elf_m", "goblin",
		"ice_zombie", "imp", "knight_f",
		"knight_m", "lizard_f", "lizard_m",
		"necromancer",
	}
	id := uuid.NewV4().String()
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	p.Unit = units.Unit{
		ID:           id,
		Action:       units.ActionIdle,
		MyCoordinate: maps.CreateCoordinate(320/2, 240/2),
		Frame:        rnd.Intn(4),
		SpriteName:   skins[rnd.Intn(len(skins))],
	}
	return p
}

func (p *Player) CreateInUnits(un units.Unit) units.IsUnit {
	p.Unit = un
	return p
}

func (p *Player) GET() *units.Unit {
	return &p.Unit
}

func (p *Player) GetID() string {
	return p.Unit.ID
}

func (p *Player) GetCoordinate() (float64, float64) {
	return p.Unit.MyCoordinate.X, p.Unit.MyCoordinate.Y
}

func (p *Player) UpdateCoordinate(x, y int) {
	p.Unit.MyCoordinate.X += float64(x)
	p.Unit.MyCoordinate.Y += float64(y)
}

func (p *Player) GetAction() string {
	return p.Unit.Action
}

func (p *Player) UpdateAction(action string) {
	log.Println(action)
	p.Action = action
	log.Println(action)

}

func (p *Player) GetRoad() *maps.Coordinate {
	return p.Unit.Road
}

func (p *Player) UpdateRoad(road *maps.Coordinate) {
	p.Road = road
}

func (p *Player) Unmarshal(b []byte) (map[string]units.IsUnit, error) {
	v := map[string]Player{}
	err := json.Unmarshal(b, &v)
	u := map[string]units.IsUnit{}
	for s, vr := range v {
		u[s] = &vr
	}
	return u, err
}
