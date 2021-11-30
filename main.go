package main

import (
	"example/my-game/actor/player"
	"example/my-game/geometry/shapes"
	"example/my-game/geometry/vector"
	"example/my-game/tiles"
	"fmt"
	"image/color"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var thePlayer *player.Player
var aTile tiles.Tile
var invisibleQuad shapes.Quad

func init() {

	thePlayer = player.InstantiatePlayer()

	X := 100.0
	Y := 250.0

	aTile = tiles.MakeTile(tiles.IndustrialTile_52, X, Y)

	var horizontalOffset float64 = 16
	var verticalOffset float64 = 16

	topLeft := vector.Vector{X: (X - horizontalOffset), Y: (Y - verticalOffset)}

	topRight := vector.Vector{X: (X + horizontalOffset), Y: (Y - verticalOffset)}

	bottomLeft := vector.Vector{X: (X - horizontalOffset), Y: (Y + verticalOffset)}

	bottomRight := vector.Vector{X: (X + horizontalOffset), Y: (Y + verticalOffset)}

	invisibleQuad = shapes.Quad{
		TopLeft:     topLeft,
		TopRight:    topRight,
		BottomLeft:  bottomLeft,
		BottomRight: bottomRight,
	}

	spew.Dump(invisibleQuad)
}

type Game struct {
	currentGameFrame int
	keys             []ebiten.Key
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	g.currentGameFrame++
	thePlayer.Update(invisibleQuad)
	return nil
}

const (
	screenWidth  = 320
	screenHeight = 240
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x00, 0x40, 0x80, 0xff})
	thePlayer.Actor.Draw(screen, g.currentGameFrame)
	aTile.Draw(screen, g.currentGameFrame)

	// keyStrs := []string{}
	// for _, p := range g.keys {
	// 	keyStrs = append(keyStrs, p.String())
	// }
	//numroutines := len(thePlayer.Actor.SubRoutines)
	//ebitenutil.DebugPrint(screen, "num routines "+strconv.Itoa(numroutines))

	// theplayersBox := thePlayer.Actor.GetHitBoxAt(thePlayer.Actor.Position)

	mx, my := ebiten.CursorPosition()
	msg := fmt.Sprintf("cursor (%d, %d)", mx, my)

	ebitenutil.DebugPrint(screen, msg+"velocityY: "+fmt.Sprint(thePlayer.Actor.Velocity.Y))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(1280, 960)
	ebiten.SetWindowTitle("GoGame")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
