package main

import (
	"fmt"
	"log"
	"strings"

	"example/my-game/actor"
	"example/my-game/actor/player"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var thePlayer *actor.Actor

func init() {
	thePlayer = player.InstantiatePlayer()
}

type Game struct {
	currentGameFrame int
	keys             []ebiten.Key
}

func (g *Game) Update() error {
	fmt.Println("keys before ", g.keys)
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	fmt.Println("keys after ", g.keys)
	g.currentGameFrame++
	return nil
}

const (
	screenWidth  = 320
	screenHeight = 240
)

func (g *Game) Draw(screen *ebiten.Image) {
	thePlayer.Draw(screen, g.currentGameFrame)

	keyStrs := []string{}
	for _, p := range g.keys {
		keyStrs = append(keyStrs, p.String())
	}
	ebitenutil.DebugPrint(screen, strings.Join(keyStrs, ", "))
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
