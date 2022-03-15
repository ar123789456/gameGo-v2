package input

import (
	"rpg/game"
	"rpg/game/maps"
	"rpg/game/units"

	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
)

//HandleMousePress handler Mouse pressed
func HandleMousePress(c *websocket.Conn, world *game.World) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		world.Players[world.MyID].UpdateRoad(maps.CreateCoordinate(ebiten.CursorPosition()))
		c.WriteJSON(
			game.Event{
				Type: game.EventTypeMove,
				Data: game.EventMove{
					UnitID:    world.MyID,
					Direction: units.DirectionLeft,
				},
			})

		return
	}
}
