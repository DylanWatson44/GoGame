package main

import (
	"log"
	"strconv"

	"example/my-game/actor/player"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var thePlayer *player.Player

func init() {
	thePlayer = player.InstantiatePlayer()
}

type Game struct {
	currentGameFrame int
	keys             []ebiten.Key
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	g.currentGameFrame++
	thePlayer.Update()
	return nil
}

const (
	screenWidth  = 320
	screenHeight = 240
)

func (g *Game) Draw(screen *ebiten.Image) {
	thePlayer.Actor.Draw(screen, g.currentGameFrame)

	keyStrs := []string{}
	for _, p := range g.keys {
		keyStrs = append(keyStrs, p.String())
	}
	numroutines := len(thePlayer.Actor.SubRoutines)
	ebitenutil.DebugPrint(screen, "num routines "+strconv.Itoa(numroutines))
	//ebitenutil.DebugPrint(screen, strings.Join(keyStrs, ", "))
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
