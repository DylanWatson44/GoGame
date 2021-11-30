package tiles

import (
	"example/my-game/geometry/vector"
	"example/my-game/sprite"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	CurrentSprite *sprite.Sprite
	Position      vector.Vector
}

type TileName string

const (
	IndustrialTile_52 TileName = "IndustrialTile_52"
)

func MakeTile(tileName TileName, X, Y float64) Tile {
	tilePath := "Tiles/" + tileName + ".png"
	tileSprite := sprite.NewSprite(string(tilePath), 32, 32, 1, 1, 0, 0, 1)
	log.Println("check ", string(tilePath))

	var tile Tile = Tile{
		CurrentSprite: tileSprite,
		Position: vector.Vector{
			X: X,
			Y: Y,
		},
	}

	return tile
}

func (tile *Tile) Draw(screen *ebiten.Image, currentGameFrame int) {
	options := &ebiten.DrawImageOptions{}
	// if actor.Facing == Left {
	// 	options.GeoM.Scale(-1, 1)
	// }
	options.GeoM.Translate(-float64(tile.CurrentSprite.FrameWidth)/2, -float64(tile.CurrentSprite.FrameHeight)/2)
	options.GeoM.Translate(tile.Position.X, tile.Position.Y)

	var imgObj *ebiten.Image = tile.CurrentSprite.GetFrame(currentGameFrame)

	screen.DrawImage(imgObj, options)
}
