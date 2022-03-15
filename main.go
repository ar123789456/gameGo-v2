package main

import (
	"fmt"
	"log"
	"rpg/game"
	"rpg/game/player"

	// "rpg/game/units"
	"rpg/input"
	"rpg/rander"

	_ "image/png"

	"github.com/gorilla/websocket"
	ebiten "github.com/hajimehoshi/ebiten/v2"
)

var world game.World
var frame int

// var img *ebiten.Image
var c *websocket.Conn

func init() {
	world = game.World{
		IsServer: false,
		// Units:    units.Units{},
		Players: player.Players{},
	}
}

// Game implements ebiten.Game interface.
type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	input.HandleKeyPress(c, &world)
	input.HandleMousePress(c, &world)
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	var randerInfo rander.Rander
	frame++
	fmt.Println(world.Players)
	randerInfo = rander.RanderBuild(&world, frame, screen)
	fmt.Println(world.Players)

	if world.Maps == nil {
		randerInfo.RanderMaps()
	}
	randerInfo.RanderUnits()

}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	var err error

	c, _, err = websocket.DefaultDialer.Dial("ws://127.0.0.1:8080/ws", nil)

	if err != nil {
		log.Println(err)
		return
	}

	go func(c *websocket.Conn) {
		defer c.Close()

		for {
			var event game.Event
			c.ReadJSON(&event)
			world.HandleEvent(&event)
		}
	}(c)

	game := &Game{}

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 440)
	ebiten.SetWindowTitle("Your game's title")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
