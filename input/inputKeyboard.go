package input

import (
	"rpg/game"
	"rpg/game/units"

	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
)

//HandleKeyPress handler key pressed
func HandleKeyPress(c *websocket.Conn, world *game.World) {
	// Write your game's logical update.

	var crEvent createEventKey

	crEvent.c = c
	crEvent.world = world

	//handler left
	crEvent.keys = []ebiten.Key{ebiten.KeyA, ebiten.KeyLeft}
	crEvent.dir = units.DirectionLeft
	if crEvent.createEventMove() {
		return
	}

	//handler right
	crEvent.keys = []ebiten.Key{ebiten.KeyD, ebiten.KeyRight}
	crEvent.dir = units.DirectionRight
	if crEvent.createEventMove() {
		return
	}

	//handler Up
	crEvent.keys = []ebiten.Key{ebiten.KeyW, ebiten.KeyUp}
	crEvent.dir = units.DirectionUp
	if crEvent.createEventMove() {
		return
	}

	//handler Down
	crEvent.keys = []ebiten.Key{ebiten.KeyS, ebiten.KeyDown}
	crEvent.dir = units.DirectionDown

	if crEvent.createEventMove() {
		return
	}

	//handle Idle
	if world.Players[world.MyID].GetAction() == units.ActionRun && world.Players[world.MyID].GetRoad() == nil {
		c.WriteJSON(game.Event{
			Type: game.EventTypeIdle,
			Data: game.EventIdle{
				UnitID: world.MyID,
			},
		})
	}
}

type createEventKey struct {
	c     *websocket.Conn
	world *game.World
	keys  []ebiten.Key
	dir   units.DirectionType
}

func (cr *createEventKey) createEventMove() bool {
	for _, key := range cr.keys {
		if ebiten.IsKeyPressed(key) {
			cr.world.Players[cr.world.MyID].UpdateRoad(nil)
			cr.c.WriteJSON(
				game.Event{
					Type: game.EventTypeMove,
					Data: game.EventMove{
						UnitID:    cr.world.MyID,
						Direction: cr.dir,
					},
				})

			return true
		}
	}
	return false
}
