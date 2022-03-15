package rander

import (
	"fmt"
	"rpg/game"
	"rpg/game/maps"
	"rpg/game/units"
	"sort"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Rander struct {
	World               *game.World
	Frame               int
	Screen              *ebiten.Image
	RelativeDifferenceX float64
	RelativeDifferenceY float64
}

func RanderBuild(world *game.World, frame int, screen *ebiten.Image) Rander {
	var r Rander
	r.World = world
	r.Frame = frame
	r.Screen = screen
	return r
}

func (rander *Rander) RanderUnits() {
	var myUnit units.IsUnit

	screenX, screenY := rander.Screen.Size()
	unitList := []units.IsUnit{}

	for uuid, unit := range rander.World.Players {
		fmt.Println(rander.World.MyID)
		fmt.Println(unit)
		if uuid == "" {
			continue
		}
		unitList = append(unitList, unit)
		// fmt.Println(unit)
		if unit.GetID() == rander.World.MyID {
			myUnit = unit
			continue
		}
	}

	sort.Slice(unitList, func(i, j int) bool {
		_, y1 := unitList[i].GetCoordinate()
		_, y2 := unitList[j].GetCoordinate()
		return y1 < y2
	})
	x, y := myUnit.GetCoordinate()
	rander.RelativeDifferenceX = float64(screenX)/2 - x
	rander.RelativeDifferenceY = float64(screenY)/2 - y

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(rander.RelativeDifferenceX, rander.RelativeDifferenceY)
	rander.Screen.DrawImage(rander.World.Maps, op)

	for _, unit := range unitList {
		un := unit.GET()
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(rander.RelativeDifferenceX+un.MyCoordinate.X, rander.RelativeDifferenceY+un.MyCoordinate.Y)
		spriteIndex := (rander.Frame/7 + un.Frame) % 4
		img, _, _ := ebitenutil.NewImageFromFile("frames/" +
			un.SpriteName + "_" + un.Action + "_anim_f" + strconv.Itoa(spriteIndex) + ".png")
		// if unit.Action == game.ActionRun {
		// 	if unit.HorizontalDirection == game.DirectionLeft {
		// 		op.GeoM.Invert()
		// 	}
		// }
		rander.Screen.DrawImage(img, op)
		// ebiten.SetWindowPosition(int(unit.X), int(unit.Y))

	}
}

func (rander *Rander) RanderMaps() {
	var x, y float64
	maps := maps.LoadMap()
	mapsImage := ebiten.NewImage(len(maps[0])*16, len(maps)*16)
	for _, line := range maps {
		for _, sprite := range line {

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(x, y)

			img, _, _ := ebitenutil.NewImageFromFile("frames/" + sprite + ".png")
			mapsImage.DrawImage(img, op)
			x += 16
		}
		y += 16
		x = 0
	}
	rander.World.Maps = mapsImage
}
